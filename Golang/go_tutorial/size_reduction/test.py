"""test docker http interface."""

import json
import requests


def test():
    """test http://127.0.0.1:5000."""
    response = requests.post("http://127.0.0.1:5000", json.dumps({"test": "test"}))
    print(response)
    print(json.loads(response.text))

if __name__ == '__main__':
    test()
