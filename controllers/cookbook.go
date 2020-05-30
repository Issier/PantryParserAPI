package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Issier/PantryParserAPI/dao"
	"github.com/Issier/PantryParserAPI/models"
)

// CookBookRootHandler routes incoming requests from cookbook to appropriate methods
func CookBookRootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(dao.GetCookbook())
	}
}

// CookBookAddHandler handles the endpoint for adding new recipes to the cookbook
func CookBookAddHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		recipe, err := parseRecipeFromPostBody(r)
		if err != nil {
			return
		}
		dao.SaveRecipe(recipe)
		json.NewEncoder(w).Encode(recipe)
	}
}

// CookBookGetRecipeHandler handles retrieving an individual recipe
func CookBookGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		recipe := dao.GetRecipe(r.URL.Path[len("/recipes/"):])
		json.NewEncoder(w).Encode(recipe)
	}
}

func parseRecipeFromPostBody(r *http.Request) (models.Recipe, error) {
	bodyDecoder := json.NewDecoder(r.Body)
	var returnVal models.Recipe
	err := bodyDecoder.Decode(&returnVal)
	if err != nil {
		return models.Recipe{}, err
	}
	return returnVal, nil
}
