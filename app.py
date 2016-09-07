from flask import Flask
app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello World!"

@app.route("/testo")
def hellos():
    return "Hello World! test"


if __name__ == "__main__":
    app.run()