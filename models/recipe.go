package models

// Recipe contains information for a recipe (name, ingredients, description)
type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
	Description string       `json:"description"`
	Link        string       `json:"link"`
}

// Ingredient defines the characteristics of an ingredient in a recipe
type Ingredient struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}
