package main

import (
	"log"
	"net/http"

	"github.com/Issier/PantryParserAPI/controllers"
)

func main() {
	controllers.SetupRecipeHandlers()
	controllers.SetupIngredientHandlers()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
