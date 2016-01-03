package main

import (
	"log"
	"net/http"
	"path"
	"html/template"
)

var dba DatabaseAccessor
var sm SessionManager

func main() {
	// initialize http server
	resourceFiles := http.FileServer(http.Dir("../templates/"))
	http.HandleFunc("/", ServeLogin)
	http.HandleFunc("/user/", ServeUser)
	http.HandleFunc("/login/", PostLogin)
	http.Handle("/js/", resourceFiles)
	http.Handle("/css/", resourceFiles)

	// initialize database connection
	err := dba.Connect("root", "", "test")
	defer dba.Close()
	if (err != nil) {
		log.Printf("SQL ERROR: %s\n", err)
		return
	}

	// initialize session manager
	sm.Init()

	// start server
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	ServeUserPage(w, "RICHARD")
	return

	cookie, _ := r.Cookie("session")
	if (cookie != nil) {
		// check cookie
		username, result := sm.SessionExists(cookie.Value)

		log.Printf("username: %s\n", username)
		// go to main page with user
		if (result) {
			ServeUserPage(w, username)
		}
	}

	errorMsg := r.URL.Query().Get("error")
	ServeLoginPage(w, errorMsg)
}

func ServeUser(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	if (cookie != nil) {
		// check cookie
		username, result := sm.SessionExists(cookie.Value)

		log.Printf("username: %s\n", username)
		// go to main page with user
		if (result) {
			// pass in username
		}
	}
}

type User struct {
	Username string
	Balance float32
}

// only call this function if user is already validated
func ServeUserPage(w http.ResponseWriter, username string) {
	file := path.Join("../templates", "user.html")

	tmpl, err := template.ParseFiles(file)
	if (err != nil) {
		log.Printf("Error Message: %s\n", err)
		return
	}

	var mainUser User
	mainUser.Username = username
	_, mainUser.Balance, _ = dba.GetUser(username)


	var allUsers []User

	userMap, _ := dba.ListUsers()
	for key, val := range userMap {
		allUsers = append(allUsers, User{key, val})
	}

	UserPageData := struct {
		MainUser User
		AllUsers []User
	} {
		mainUser,
		allUsers,
	}

	tmpl.Execute(w, UserPageData)
}

func ServeLoginPage(w http.ResponseWriter, errorMsg string) {
	file := path.Join("../templates", "login.html")

	tmpl, err := template.ParseFiles(file)
	if (err != nil) {
		log.Printf("Error Message: %s\n", err)
		return
	}

	LoginPageData := struct {
		errorMessage string
	} {
		errorMsg,
	}
	tmpl.Execute(w, LoginPageData)
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	username := r.FormValue("username")
	password := r.FormValue("password")

	if (username == "") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty username\n"))
		return
	}
	
	if (password == "") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty Password\n"))
		return
	}

	checkedPassword, _, err := dba.GetUser(username)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Such Username"))
		return
	}

	if (checkedPassword != password) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Password did not match"))
		return
	}

	var cookie http.Cookie
	cookie.Name = "session"
	cookie.Value = sm.GenerateNewSessionId(username)
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

