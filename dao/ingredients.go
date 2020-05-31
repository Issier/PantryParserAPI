package dao

import (
	"github.com/Issier/PantryParserAPI/models"
)

var knownIngredients map[string]models.Ingredient

// GetIngredients return all known ingredients
func GetIngredients() map[string]models.Ingredient {
	return knownIngredients
}

func init() {
	knownIngredients = map[string]models.Ingredient{}
	for _, recipe := range GetCookbook() {
		for _, ingredient := range recipe.Ingredients {
			knownIngredients[ingredient.Name] = ingredient
		}
	}
}
