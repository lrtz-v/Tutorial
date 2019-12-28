"""test python app run in docker."""

import logging
from flask import Flask, request
import json
app = Flask(__name__)

logging.basicConfig(filename='./logs/log.log', level=logging.INFO)


@app.route('/', methods=["POST"])
def hello_world():
    """hello_world test python app run in docker."""
    data = request.get_data()
    logging.info(data)
    return json.dumps({"message": "success", "code": 200})

if __name__ == '__main__':
    app.run(host="0.0.0.0", port=80)
