### Overview
The Fine Ant Finance App! Brought to you by Alectryon Tech-tryon, a simple web-app to help you manage you finances among a group of people, great for college and roommates!


### Requirements
It's probably best to put all all of the environment stuff into a virtual env, but w/e
* Python + Python Flask: "*pip install flask*"
* Flask My-SQL extension: "pip install flask-mysql"
* MySQL/MariaDB: "sudo apt-get install mysql-server" or "brew install mariadb" 
* GUnicorn: run the server or something

### How to Configure and Run
* Install the appropriate libraries
* Properly Configure the Database, make sure it's running
* Run application: "*python app.py*"
* Alternatively, you can serve it with Gunicorn

After we finish the project, we'll add a couple of init scripts.

### Structure
*Front-End*: Using the standard web-kit (html, css, javascript + jQuery), a basic interface that does exactly what you think it does, it sends the requests over to the Back-End for processing.

*Back End*: The Flask Layer receives requests from the front-end via get/post requests and processes them.

*MySQL*: Used MariaDB on Mac because it's way easier to install. in the scripts folder, There'll be scripts to various things such as load/reload the SQL Tables, for the sake of example, we're going to use:
* User: "finance"
* Password: "password"

#### Libraries/Things Used
* Python Flask
* MySQL
* Gunicorn
* Bootstrap
* [picitelli's simple modal library](https://github.com/picitelli/js-modal)
