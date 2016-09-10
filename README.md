### Structure and Overview
Front End is written with stand web things (html, css, javascript)

Back End is written with Python and Flask, database is MySQL


### Requirements
It's probably best to put all all of the environment stuff into a virtual env, but w/e
* Python + Python Flask: "*pip install flask*"
* Flask My-SQL extension: "pip install flask-mysql"
* MySQL/MariaDB: "sudo apt-get install mysql-server" or "brew install mariadb" 
* GUnicorn: run the server or something

MySQL will have to configured during the install, if MySQL is too clunky, use

### How to Run
* Install the appropriate libraries


Command runs it on localhost of the machine
* "*python app.py*"

### Structure

*MySQL*: Used MariaDB on Mac because it's way easier to install. in the scripts folder, There'll be scripts to various things such as load/reload the SQL Tables, for the sake of example, we're going to use:
* User: "finance"
* Password: "Password"
That will create all the tables.


#### Libraries Used
* Python Flask
* [picitelli's simple modal library](https://github.com/picitelli/js-modal)
