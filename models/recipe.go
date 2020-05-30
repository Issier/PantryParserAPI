package models

// Recipe contains information for a recipe (name, ingredients, description)
type Recipe struct {
	Name        string
	Ingredients []Ingredient
	Description string
}

// Ingredient defines the characteristics of an ingredient in a recipe
type Ingredient struct {
	Name     string
	Category string
}
