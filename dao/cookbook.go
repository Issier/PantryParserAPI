package dao

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/Issier/PantryParserAPI/models"
	_ "github.com/go-sql-driver/mysql"
)

var cookBook map[string]models.Recipe
var cookBookLock = sync.RWMutex{}
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
	db, err := sql.Open("mysql", config.DBString)
	defer db.Close()
	if err != nil {
		panic("Unable to connect to database")
	}
	conn, err := db.Begin()
	if err != nil {
		return errors.New("Unable to establish a transaction")
	}
	result, err := conn.Exec("insert into recipe (recipe_name, recipe_description) values (?, ?)", recipe.Name, recipe.Description)
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
		ingredientRow := conn.QueryRow("select id from ingredient where ingredient_name=?", ingredient.Name)
		var ingredientID int
		ingredientRow.Scan(&ingredientID)
		_, err = conn.Exec("insert into cookbookentry (recipe_id, ingredient_id) values (?, ?)", recipeID, ingredientID)
		if err != nil {
			conn.Rollback()
			return errors.New("Recipe entry already exists")
		}
	}
	conn.Commit()
	return nil
}

// GetRecipe retrieves a recipe by the given key
func GetRecipe(name string) models.Recipe {
	cookBookLock.RLock()
	defer cookBookLock.RUnlock()
	return cookBook[name]
}

// GetCookbook returns the entire cookbook
func GetCookbook() map[string]models.Recipe {
	cookBookLock.RLock()
	defer cookBookLock.RUnlock()
	return cookBook
}

func init() {
	cookBook = map[string]models.Recipe{
		"Grilled Cheese": {
			Name: "Grilled Cheese",
			Ingredients: []models.Ingredient{
				{
					Name:     "Cheese",
					Category: "Dairy",
				},
				{
					Name:     "Bread",
					Category: "Grain",
				},
				{
					Name:     "Butter",
					Category: "Dairy",
				},
			},
			Description: "Grilled cheese so delicious you'll want to eat it",
		},
		"Pan Steak": {
			Name: "Pan Steak",
			Ingredients: []models.Ingredient{
				{
					Name:     "Ribeye",
					Category: "Meat",
				},
				{
					Name:     "Thyme",
					Category: "Herb",
				},
				{
					Name:     "Butter",
					Category: "Dairy",
				},
			},
			Description: "Delicious pan seared steak",
		},
		"Quick Chickpea and Spinach Stew": {
			Name:        "Quick Chickpea and Spinach Stew",
			Description: "A delicious vegetarian stew that will last for days",
			Ingredients: []models.Ingredient{
				{
					Name:     "Can of Whole Tomatoes",
					Category: "Fruit",
				},
				{
					Name:     "Ginger",
					Category: "Spice",
				},
				{
					Name:     "Extra Virgin Olive Oil",
					Category: "Oil",
				},
				{
					Name:     "Medium Onion",
					Category: "Vegetable",
				},
				{
					Name:     "Garlic (cloves)",
					Category: "Vegetable",
				},
				{
					Name:     "Smoked Paprika",
					Category: "Spice",
				},
				{
					Name:     "Spinach",
					Category: "Vegetable",
				},
				{
					Name:     "Chickpeas",
					Category: "Beans",
				},
				{
					Name:     "Bay Leaves",
					Category: "Spice",
				},
				{
					Name:     "Soy Sauce",
					Category: "Sauce",
				},
				{
					Name:     "Sherry Vinegar",
					Category: "Vinegar",
				},
			},
		},
	}
}
