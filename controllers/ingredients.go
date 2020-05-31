package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Issier/PantryParserAPI/dao"
)

// GetIngredientsHandler returns a list of recognized ingredients
func GetIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(dao.GetIngredients())
}
