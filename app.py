#Python Libs and Imports
# from flask import Flask
from flask import *
#Importing the blueprint object from other Files
from login import login
from user import user

#Sets the application Variable name
app = Flask(__name__)

#Register routes in other files
app.register_blueprint(login)
app.register_blueprint(user)



#This is an app route, a route is pretty much one URL
@app.route("/")
def hello():
    return "Hello World!" #what the route returns is the data that will be rendered onto the page

# As you can see, we can have many different Routes. We'll want to split
# this into multiple files though, which we do by registering routes from other files
# which you can see above.
@app.route("/testo")
def hellos():
    return "Something Else"


if __name__ == "__main__":
    app.run()