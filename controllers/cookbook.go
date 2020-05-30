package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Issier/PantryParserAPI/dao"
	"github.com/Issier/PantryParserAPI/models"
)

// CookBookRootHandler routes incoming requests from cookbook to appropriate methods
func CookBookRootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if len(r.URL.Query().Get("ingredients")) > 0 {
			cookBookGetRecipeByIngredients(w, r)
		} else {
			json.NewEncoder(w).Encode(dao.GetCookbook())
		}
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

// CookBookGetRecipeByIngredients returns a map of number of matching ingredients to arrays of recipes
func cookBookGetRecipeByIngredients(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		ingredients := r.URL.Query().Get("ingredients")
		recipesIncludingContent := map[int][]models.Recipe{}
		cookBook := dao.GetCookbook()
		for _, recipe := range cookBook {
			matchingIngredients := 0
			for _, ingredient := range recipe.Ingredients {
				if strings.Contains(ingredients, ingredient.Name) {
					matchingIngredients++
				}
			}
			if matchingIngredients > 0 {
				recipesIncludingContent[matchingIngredients] = append(recipesIncludingContent[matchingIngredients], recipe)
			}
		}
		json.NewEncoder(w).Encode(recipesIncludingContent)
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
