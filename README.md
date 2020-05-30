# PantryParserAPI
The backend for [Pantry Parser](https://github.com/Issier/PantryParser)

## Current Endpoints

### GET: /recipes
Returns all recipes in current Cookbook

### POST: /recipes/add
Stores recipe provided in POST body

### GET: /recipes/{recipe name}
Returns individual recipe by name, returning an empty recipe if not found in cookbook
