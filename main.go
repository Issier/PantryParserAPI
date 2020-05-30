package main

import (
	"net/http"

	"github.com/Issier/PantryParserAPI/controllers"
)

func main() {
	http.HandleFunc("/recipes", controllers.CookBookRootHandler)
	http.HandleFunc("/recipes/add", controllers.CookBookAddHandler)
	http.HandleFunc("/recipes/", controllers.CookBookGetRecipeHandler)

	http.ListenAndServe(":3000", nil)
}
