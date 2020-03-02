import os
from flask import Flask, request
from flask_restful import Resource, Api
import requests

app = Flask(__name__)
api = Api(app)

cache = {}

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
        if cache.get(ingredients):
            return cache[ingredients]
        else:
            recipes = requests.get('https://api.spoonacular.com/recipes/findByIngredients', params={'ingredients': ingredients, 'apiKey': self.key}).json()
            cache[ingredients] = recipes
            return recipes

api.add_resource(RecipeByIngredientSimpleCache, '/')

if __name__ == '__main__':
    app.run()