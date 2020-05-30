package dao

import (
	"github.com/Issier/PantryParserAPI/models"
)

var cookBook map[string]models.Recipe

// SaveRecipe persists provided recipe
func SaveRecipe(recipe models.Recipe) error {
	cookBook[recipe.Name] = recipe
	return nil
}

// GetRecipe retrieves a recipe by the given key
func GetRecipe(name string) models.Recipe {
	return cookBook[name]
}

// GetCookbook returns the entire cookbook
func GetCookbook() map[string]models.Recipe {
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
			Name: "Quick Chickpea and Spinach Stew",
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
