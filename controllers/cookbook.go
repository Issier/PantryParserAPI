package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Issier/PantryParserAPI/cors"
	"github.com/Issier/PantryParserAPI/dao"
	"github.com/Issier/PantryParserAPI/models"
)

func SetupRecipeHandlers() {
	http.Handle("/recipes", cors.CorsHandler(http.HandlerFunc(CookBookRootHandler)))
	http.Handle("/recipes/add", cors.CorsHandler(http.HandlerFunc(CookBookAddHandler)))
	http.Handle("/recipes/", cors.CorsHandler(http.HandlerFunc(CookBookGetRecipeHandler)))
}

// CookBookRootHandler routes incoming requests from cookbook to appropriate methods
func CookBookRootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		encoder := json.NewEncoder(w)
		recipesGroupedByIngredientMatchCount, _ := dao.GetRecipesByIngredients(strings.Split(r.URL.Query().Get("ingredients"), ","))
		encoder.Encode(recipesGroupedByIngredientMatchCount)
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
		err = dao.SaveRecipe(recipe)
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(recipe)
	}
}

// CookBookGetRecipeHandler handles retrieving an individual recipe
func CookBookGetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		recipe, err := dao.GetRecipe(r.URL.Path[len("/recipes/"):])
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		} else if recipe.Name == "" {
			w.WriteHeader(http.StatusNotFound)
		}
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
