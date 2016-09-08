from flask import Blueprint
from flask import Flask, request, render_template

user = Blueprint('user', __name__)

@user.route("/user")
def renderText():
    return render_template("user.html")