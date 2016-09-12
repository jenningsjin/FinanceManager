#Python Libs and Imports
from flask import Flask
from flaskext.mysql import MySQL
from flask import render_template

#Importing the blueprint object from other Files
from login import login
from user import user

#Sets the application Variable name
app = Flask(__name__)


# MySQL configurations, pretty self explanatory
# This stuff also needs to be 
mysql = MySQL()
app.config['MYSQL_DATABASE_USER'] = 'finance'
app.config['MYSQL_DATABASE_PASSWORD'] = 'password'
app.config['MYSQL_DATABASE_DB'] = 'FINANCE_APP'
app.config['MYSQL_DATABASE_HOST'] = 'localhost'
mysql.init_app(app)

# 2) Register routes in other files
# 	 this is python so you have to import everything, then we register
#	 all the routes from other files so we can access them in the app
#	 All the other imports and stuff are above, and are pretty self-explanatory
app.register_blueprint(login)
app.register_blueprint(user)
#MySQL Database Configs



## 1) Start reading here, go down until prompted to move to a different section
#	  This is an app route, a route is pretty much one URL, we've defined the root address
#	  below to return a page that just says "hello world"
@app.route("/")
def hello():
    return "Hello World!" #what the route returns is the data that will be rendered onto the page

# 1.1) As you can see, we can have many different Routes. We'll want to split
# this into multiple files though, which we do by registering routes from other files
# which you can see above at 2)
@app.route("/testo")
def hellos():
	# So in more detail, what happens is that when we go to https://<root-address>/testo,
	# this route testo, is executed, so for our intents and purposes, the function hellos gets run
	# Whatever we return is rendered to the page. So this is the spot where we can do preprocessing
	# of whatever data we pull, which we can then toss into the template

	data = {  #we've defined a good ol' python dictionary here, but in my mind, it's really a json by the
		'example': "cheese",  # time we get to the js side.
		'number': 10
	}
	return render_template('test.html', data=data)
	# We can see above that when we return a function that says renders a template (literally named)
	# if we have a "templates" folder, then flask will automatically look there.
	#
	# Also, we pass in any data that we want to access in the template
	# into the function. I believe you can pass in as many variables as you want, but it's way easier than
	# that to just pass in an object, it's defined as: <name in template> = <variable you pass in>
	# the template name can be anything

# 1) is over, go to 2) above
if __name__ == "__main__":
    app.run()

#root@localhost: k9upiI;TID!>
