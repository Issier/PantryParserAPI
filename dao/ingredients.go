package dao

import (
	"database/sql"
	"errors"

	"github.com/Issier/PantryParserAPI/models"
)

// GetIngredients return all known ingredients
func GetIngredients() ([]models.Ingredient, error) {
	return getIngredients("")
}

// GetIngredientsByRecipeName returns ingredients used in provided recipe
func GetIngredientsByRecipeName(name string) ([]models.Ingredient, error) {
	return getIngredients(name)
}

func getIngredients(name string) ([]models.Ingredient, error) {
	conn, err := sql.Open("mysql", config.DBString)
	defer conn.Close()
	if err != nil {
		return []models.Ingredient{}, errors.New("Unable to begin session")
	}

	var rows *sql.Rows
	if name == "" {
		rows, _ = conn.Query("select ingredient_name, ingredient_category from ingredient")
	} else {
		rows, _ = conn.Query("select ingredient_name, ingredient_category from cookbookentry inner join ingredient on ingredient_id = ingredient.id inner join recipe on recipe_id = recipe.id  where recipe_name = ?", name)
	}

	ingredients := []models.Ingredient{}
	for rows.Next() {
		ingredient := models.Ingredient{}
		rows.Scan(&ingredient.Name, &ingredient.Category)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}
