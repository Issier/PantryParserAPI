package main

import (
	"net/http"

	"github.com/Issier/PantryParserAPI/controllers"
)

func main() {
	http.HandleFunc("/recipes", controllers.CookBookRootHandler)
	http.HandleFunc("/recipes/add", controllers.CookBookAddHandler)
	http.HandleFunc("/recipes/", controllers.CookBookGetRecipeHandler)
	http.HandleFunc("/ingredients", controllers.GetIngredientsHandler)

	http.ListenAndServe(":8080", nil)
}
