package dao

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Issier/PantryParserAPI/models"
	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", config.DBString)
	defer db.Close()
	if err != nil {
		panic("Unable to connect to database")
	}
	conn, err := db.Begin()
	if err != nil {
		return errors.New("Unable to establish a transaction")
	}
	result, err := conn.Exec("insert into recipe (recipe_name, recipe_description) values ($1, $2)", recipe.Name, recipe.Description)
	if err != nil {
		conn.Rollback()
		return errors.New("Recipe with that name already exists")
	}

	recipeID, _ := result.LastInsertId()

	if err != nil {
		conn.Rollback()
		return errors.New("Unable to query database")
	}

	for _, ingredient := range recipe.Ingredients {
		ingredientRow := conn.QueryRow("select id from ingredient where ingredient_name=$1", ingredient.Name)
		var ingredientID int
		ingredientRow.Scan(&ingredientID)
		_, err = conn.Exec("insert into cookbookentry (recipe_id, ingredient_id) values ($1, $2)", recipeID, ingredientID)
		if err != nil {
			conn.Rollback()
			return errors.New("Recipe entry already exists")
		}
	}
	conn.Commit()
	return nil
}

// GetRecipe retrieves a recipe by the given key
func GetRecipe(name string) (models.Recipe, error) {
	conn, err := sql.Open("postgres", config.DBString)
	defer conn.Close()
	if err != nil {
		return models.Recipe{}, errors.New("Unable to begin session")
	}
	recipeRows, err := conn.Query("select ingredient_name, ingredient_category, recipe_name, recipe_description from cookbookentry inner join recipe on recipe_id = recipe.id inner join ingredient on ingredient_id = ingredient.id where recipe.recipe_name = $1", name)
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
	conn, err := sql.Open("postgres", config.DBString)
	defer conn.Close()
	if err != nil || len(ingredients) == 0 {
		return map[int][]models.Recipe{}, errors.New("Unable to begin session")
	}
	matchingIngredientString := "ingredient_name = $1"
	for range ingredients[1:] {
		matchingIngredientString += " OR ingredient_name = $1"
	}
	interfaces := make([]interface{}, len(ingredients))
	for i, ingredient := range ingredients {
		interfaces[i] = ingredient
	}
	recipeRows, err := conn.Query("select recipe_name, recipe_description, COUNT(recipe_name) "+
		"from cookbookentry inner join recipe on recipe_id = recipe.id inner join ingredient "+
		"on ingredient_id = ingredient.id where "+matchingIngredientString+" group by recipe_name", interfaces...)
	if err != nil {
		return map[int][]models.Recipe{}, errors.New("Unable to pull information")
	}
	recipes := make(map[int][]models.Recipe)
	for recipeRows.Next() {
		recipe := models.Recipe{}
		var numberOccurences int
		recipeRows.Scan(&recipe.Name, &recipe.Description, &numberOccurences)
		recipe.Ingredients, _ = GetIngredientsByRecipeName(recipe.Name)
		recipes[numberOccurences] = append(recipes[numberOccurences], recipe)
	}
	return recipes, nil
}
