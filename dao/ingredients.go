package dao

import (
	"database/sql"
	"errors"

	"github.com/Issier/PantryParserAPI/models"
)

// GetIngredients return all known ingredients
func GetIngredients() ([]models.Ingredient, error) {
	conn, err := sql.Open("mysql", config.DBString)
	defer conn.Close()
	if err != nil {
		return []models.Ingredient{}, errors.New("Unable to begin session")
	}
	rows, _ := conn.Query("select ingredient_name, ingredient_category from ingredient")

	ingredients := []models.Ingredient{}
	for rows.Next() {
		ingredient := models.Ingredient{}
		rows.Scan(&ingredient.Name, &ingredient.Category)
		ingredients = append(ingredients, ingredient)
	}
	return ingredients, nil
}
