package main

import (
	"log"
	"net/http"
	"path"
	"html/template"
)

var dba DatabaseAccessor

func main() {
	jsFs := http.FileServer(http.Dir("../templates/"))
	http.HandleFunc("/", ServeLoginPage)
	http.HandleFunc("/user/", ServeUser)
	http.HandleFunc("/login/", PostLogin)
	http.Handle("/js/", jsFs)

	err := dba.Connect("root", "", "test")
	defer dba.Close()
	if (err != nil) {
		log.Printf("SQL ERROR: %s\n", err)
		return
	}

	var usersMap map[string]float32
	usersMap, err = dba.ListUsers()
	if (err != nil) {
		log.Printf("SQL ERROR: %s\n", err)
		return
	}

	for user, balance := range usersMap {
		log.Printf("User %s has balance: %f\n", user, balance)
	}


	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

type LoginPageData struct {
	Error string
}

func ServeLoginPage(w http.ResponseWriter, r *http.Request) {
	// cookie, tmpErr := r.Cookie("session")
	file := path.Join("../templates", "login.html")

	errorMsg := r.URL.Query().Get("error")

	tmpl, err := template.ParseFiles(file)
	if (err != nil) {
		log.Printf("Error Message: %s\n", err)
		return
	}

	LoginPageData := LoginPageData{Error:errorMsg}
	tmpl.Execute(w, LoginPageData)
}

func ServeUser(w http.ResponseWriter, r *http.Request) {

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

	checkedPassword, err := dba.UserPassword(username)
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
	cookie.Value = "value"
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

