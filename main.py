from flask import Flask
from flask_restful import Resource, Api

app = Flask(__name__)
api = Api(app)

class IsAlive(Resource):
    def get(self):
        return {'app': 'PantryParserAPI'}

api.add_resource(IsAlive, '/')

if __name__ == '__main__':
    app.run()