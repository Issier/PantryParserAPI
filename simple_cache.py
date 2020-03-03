import requests

class SimpleCache:

    def __init__(self):
        self.cache = {}

    def get_resource(self, url, key, params = {}):
        if self.cache.get(key):
            return self.cache[key]
        else:
            resource = requests.get(url, params = params).json()
            self.cache[key] = resource
            return resource

