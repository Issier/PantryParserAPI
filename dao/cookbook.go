package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"strconv"

	"github.com/Issier/PantryParserAPI/models"
)

var config configFile

type configFile struct {
	DBString string `json:"dbString"`
}

func init() {
	confFile, err := os.Open("conf.json")
	if err != nil {
		fmt.Println(err)
	}
	defer confFile.Close()

	decoder := json.NewDecoder(confFile)
	config = configFile{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}
}

// SaveRecipe persists provided recipe
func SaveRecipe(recipe models.Recipe) error {
	conn, err := pgxpool.Connect(context.Background(), config.DBString)
	if err != nil {
		return errors.New("unable to establish a transaction")
	}
	defer conn.Close()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return errors.New("unable to begin transaction")
	}
	defer tx.Rollback(context.Background())

	fmt.Println(recipe.Name, recipe.Description)

	_, err = tx.Exec(context.Background(), "insert into recipe (recipe_name, recipe_description) values ($1, $2)", recipe.Name, recipe.Description)
	if err != nil {
		return errors.New("recipe with that name already exists")
	}

	recipeRow := tx.QueryRow(context.Background(), "select id, recipe_name from recipe where recipe_name=$1", recipe.Name)
	var recipeID int
	var recipeName string
	recipeRow.Scan(&recipeID, &recipeName)

	for _, ingredient := range recipe.Ingredients {
		ingredientRow := conn.QueryRow(context.Background(), "select id from ingredient where ingredient_name=$1", ingredient.Name)
		var ingredientID int
		ingredientRow.Scan(&ingredientID)

		_, err = tx.Exec(context.Background(), "insert into cookbookentry (recipe_id, ingredient_id) values ($1, $2)", recipeID, ingredientID)
		if err != nil {
			return errors.New("recipe entry already exists")
		}
		if err != nil {
			return errors.New("failed to commit cook book entry")
		}
	}
	tx.Commit(context.Background())
	return nil
}

// GetRecipe retrieves a recipe by the given key
func GetRecipe(name string) (models.Recipe, error) {
	conn, err := pgxpool.Connect(context.Background(), config.DBString)
	defer conn.Close()
	if err != nil {
		return models.Recipe{}, errors.New("Unable to begin session")
	}
	recipeRows, err := conn.Query(context.Background(), "select ingredient_name, ingredient_category, recipe_name, recipe_description from cookbookentry inner join recipe on recipe_id = recipe.id inner join ingredient on ingredient_id = ingredient.id where recipe.recipe_name = $1", name)
	if err != nil {
		return models.Recipe{}, errors.New("Unable to pull information")
	}
	recipeRows.Next()
	recipe := models.Recipe{}
	ingredient := models.Ingredient{}
	recipeRows.Scan(&ingredient.Name, &ingredient.Category, &recipe.Name, &recipe.Description)
	recipe.Ingredients = append(recipe.Ingredients, ingredient)
	for recipeRows.Next() {
		recipeRows.Scan(&ingredient.Name, &ingredient.Category)
		recipe.Ingredients = append(recipe.Ingredients, ingredient)
	}
	return recipe, nil
}

func GetRecipesByIngredients(ingredients []string) (map[int][]models.Recipe, error) {
	conn, err := pgxpool.Connect(context.Background(), config.DBString)
	defer conn.Close()
	if err != nil || len(ingredients) == 0 {
		return map[int][]models.Recipe{}, errors.New("Unable to begin session")
	}
	matchingIngredientString := "ingredient_name = $1"
	for index, _ := range ingredients[1:] {
		matchingIngredientString += " OR ingredient_name = $" + strconv.Itoa(index+2)
	}
	interfaces := make([]interface{}, len(ingredients))
	for i, ingredient := range ingredients {
		interfaces[i] = ingredient
	}
	recipeRows, err := conn.Query(context.Background(), "select recipe_name, recipe_description, COUNT(recipe_name) "+
		"from cookbookentry inner join recipe on recipe_id = recipe.id inner join ingredient "+
		"on ingredient_id = ingredient.id where "+matchingIngredientString+" group by recipe_name, recipe_description", interfaces...)
	if err != nil {
		return map[int][]models.Recipe{}, errors.New("Unable to pull information")
	}
	recipes := make(map[int][]models.Recipe)
	for recipeRows.Next() {
		recipe := models.Recipe{}
		var numberOccurences int
		recipeRows.Scan(&recipe.Name, &recipe.Description, &numberOccurences)
		fmt.Println(recipe)
		recipe.Ingredients, _ = GetIngredientsByRecipeName(recipe.Name)
		recipes[numberOccurences] = append(recipes[numberOccurences], recipe)
	}
	return recipes, nil
}
