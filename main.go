package main

import (
	"net/http"

	"github.com/Issier/PantryParserAPI/controllers"
)

func main() {
	controllers.SetupRecipeHandlers()
	controllers.SetupIngredientHandlers()

	http.ListenAndServe(":8080", nil)
}
