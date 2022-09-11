package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Issier/PantryParserAPI/models"
	"net/http"

	"github.com/Issier/PantryParserAPI/cors"
	"github.com/Issier/PantryParserAPI/dao"
)

func SetupIngredientHandlers() {
	http.Handle("/ingredients", cors.CorsHandler(http.HandlerFunc(GetIngredientsHandler)))
	http.Handle("/ingredients/add", cors.CorsHandler(http.HandlerFunc(IngredientAddHandler)))
}

// GetIngredientsHandler returns a list of recognized ingredients
func GetIngredientsHandler(w http.ResponseWriter, r *http.Request) {
	ingredients, _ := dao.GetIngredients()
	json.NewEncoder(w).Encode(ingredients)
}

// IngredientAddHandler handles the endpoint for adding new ingredients
func IngredientAddHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		ingredient, err := parseIngredientFromPostBody(r)
		if err != nil {
			return
		}
		err = dao.SaveIngredient(ingredient)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(ingredient)
	}
}

func parseIngredientFromPostBody(r *http.Request) (models.Ingredient, error) {
	bodyDecoder := json.NewDecoder(r.Body)
	var returnVal models.Ingredient
	err := bodyDecoder.Decode(&returnVal)
	if err != nil {
		return models.Ingredient{}, err
	}
	return returnVal, nil
}
