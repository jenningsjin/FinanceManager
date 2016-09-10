from flask import Blueprint

login = Blueprint('login', __name__)

@login.route("/login")
def accountList():
	return "login page"

@login.route('/signUp',methods=['POST'])
def signUp():
	return "default"