# PantryParserAPI
The backend for [Pantry Parser](https://github.com/Issier/PantryParser)

## Current Endpoints

### GET: /recipes
Returns all recipes in current Cookbook

### GET: /recipes?ingredients={list}
Returns recipes matching the provided ingredients where list is a comma seperated list of ingredients. Return values are grouped by number of matching ingredients

### POST: /recipes/add
Stores recipe provided in POST body

### GET: /recipes/{recipe name}
Returns individual recipe by name, returning an empty recipe if not found in cookbook

### GET: /ingredients
Returns list of known ingredients
