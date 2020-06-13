package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Issier/PantryParserAPI/cors"
	"github.com/Issier/PantryParserAPI/dao"
)

func SetupIngredientHandlers() {
	http.Handle("/ingredients", cors.CorsHandler(http.HandlerFunc(GetIngredientsHandler)))
}

// GetIngredientsHandler returns a list of recognized ingredients
func GetIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	ingredients, _ := dao.GetIngredients()
	json.NewEncoder(w).Encode(ingredients)
}
