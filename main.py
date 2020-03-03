import os
from flask import Flask, request
from flask_restful import Resource, Api
from simple_cache import SimpleCache

app = Flask(__name__)
api = Api(app)

recipe_query_cache = SimpleCache()
recipe_cache = SimpleCache()

class RecipeByIngredientSimpleCache(Resource):
    """Returns recipes based on provided ingredients using the Spoonacular API, 
       using a simple caching mechanism to avoid multiple API calls with the same argument
    """
    def __init__(self):
        self.key = os.environ['SPOON_API_KEY']

    def get(self):
        if request.args.get('ingredients'):
            ingredients = frozenset({ingr for ingr in request.args['ingredients'].split(',')})
        else: 
            return {}
        return recipe_query_cache.get_resource(
            'https://api.spoonacular.com/recipes/findByIngredients',
            ingredients, 
            {'ingredients': request.args['ingredients'], 'apiKey': self.key}
        )


class RecipeById(Resource):
    """Returns information about a specific ingredient based on its ID"""

    def __init__(self):
        self.key = os.environ['SPOON_API_KEY']

    def get(self, recipe_id):
        return recipe_cache.get_resource(
            f"https://api.spoonacular.com/recipes/{recipe_id}/information", 
            recipe_id, 
            {'apiKey': self.key}
        )
        

api.add_resource(RecipeByIngredientSimpleCache, '/api/v1/getRecipes')
api.add_resource(RecipeById, '/api/v1/getRecipes/<int:recipe_id>')

if __name__ == '__main__':
    app.run()