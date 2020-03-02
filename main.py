import os
from flask import Flask
from flask_restful import Resource, Api
import requests

app = Flask(__name__)
api = Api(app)

class IsAlive(Resource):
    def __init__(self):
        self.key = os.environ['SPOON_API_KEY']

    def get(self):
        return requests.get('https://api.spoonacular.com/recipes/findByIngredients', params={'ingredients': 'shrimp,pasta', 'apiKey': self.key}).json()

api.add_resource(IsAlive, '/')

if __name__ == '__main__':
    app.run()