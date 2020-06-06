package dao

import (
	"sync"

	"github.com/Issier/PantryParserAPI/models"
)

var knownIngredients map[string]models.Ingredient
var ingredientsLock = sync.RWMutex{}

// GetIngredients return all known ingredients
func GetIngredients() []models.Ingredient {
	ingredientsLock.RLock()
	defer ingredientsLock.RUnlock()
	ingredients := []models.Ingredient{}
	for _, ingredient := range knownIngredients {
		ingredients = append(ingredients, ingredient)
	}
	return ingredients
}

func init() {
	knownIngredients = map[string]models.Ingredient{}
	for _, recipe := range GetCookbook() {
		for _, ingredient := range recipe.Ingredients {
			knownIngredients[ingredient.Name] = ingredient
		}
	}
}
