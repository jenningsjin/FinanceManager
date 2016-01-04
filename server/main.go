package main

import (
	"log"
	"net/http"
	"path"
	"html/template"
	"strconv"
	"sort"
)

var dba DatabaseAccessor
var sm SessionManager

func main() {
	// initialize http server
	resourceFiles := http.FileServer(http.Dir("../templates/"))
	http.HandleFunc("/", ServeLogin)
	http.HandleFunc("/user/", ServeUser)
	http.HandleFunc("/login/", PostLogin)
	http.HandleFunc("/signup/", PostSignup)
	http.HandleFunc("/logout/", PostLogout)
	http.HandleFunc("/add-transaction/", PostAddTransaction)
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
	cookie, _ := r.Cookie("session")
	if (cookie != nil) {
		// check cookie
		username, result := sm.SessionExists(cookie.Value)

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

		// go to main page with user
		if (result) {
			// pass in username
			ServeUserPage(w, username)
		} else {
			ServeLoginPage(w, "")
		}
	} else {
		log.Printf("User not validated\n")
		ServeLoginPage(w, "")
	}
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
	_, mainUser.Balance, _, _ = dba.GetUser(username)


	var allUsers []User

	userMap, _ := dba.ListUsers()
	for key, val := range userMap {
		allUsers = append(allUsers, User{key, val, 0})
	}

	var transactions []Transaction
	transactions, err = dba.GetTransactions()

	sort.Sort(UserSlice(allUsers))

	// transactions needed for everyone to be at 0 balance
	var balanceTransactions []Transaction
	var idxFront, idxBack int
	idxFront = 0
	idxBack = len(allUsers) - 1

	var frontBalance, backBalance float32
	frontBalance = Abs(allUsers[idxFront].Balance)
	backBalance = Abs(allUsers[idxBack].Balance)

	for (idxFront != idxBack) {
		if (frontBalance == 0 && backBalance == 0) {
			break
		}
		var transaction Transaction
		if (frontBalance < backBalance) {
			transaction = Transaction{allUsers[idxFront].Username,
				allUsers[idxBack].Username, frontBalance, "", ""}
			backBalance -= frontBalance
			idxFront++
			frontBalance = Abs(allUsers[idxFront].Balance)
		} else {
			transaction = Transaction{allUsers[idxFront].Username,
				allUsers[idxBack].Username, backBalance, "", ""}
			frontBalance -= backBalance
			idxBack--
			backBalance = Abs(allUsers[idxBack].Balance)
		}
		balanceTransactions = append(balanceTransactions, transaction)
	}

	UserPageData := struct {
		MainUser User
		AllUsers []User
		Transactions []Transaction
		BalanceTransactions []Transaction
	} {
		mainUser,
		allUsers,
		transactions,
		balanceTransactions,
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

	checkedPassword, _, _, err := dba.GetUser(username)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No Such Username"))
		return
	}

	if (checkedPassword != password) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Incorrect Password"))
		return
	}

	var cookie http.Cookie
	cookie.Name = "session"
	cookie.Value = sm.GenerateNewSessionId(username)
	cookie.Domain = ""
	cookie.Path = "/"
	cookie.MaxAge = 0

	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func PostSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "test/plain")

	username := r.FormValue("username")
	password := r.FormValue("password")

	err := dba.CreateUser(username, password)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}

func PostLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	cookie, _ := r.Cookie("session")
	sm.DeleteSession(cookie.Value)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}

func PostAddTransaction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	debtor := r.FormValue("debtor")
	debtee := r.FormValue("debtee")
	amount := r.FormValue("amount")
	description := r.FormValue("description")

	var amountFloat float32
	var amountFloat64 float64
	var floatErr error
	amountFloat64, floatErr = strconv.ParseFloat(amount, 32)
	if (floatErr != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(floatErr.Error()))
	}
	amountFloat = float32(amountFloat64)

	err := dba.CreateTransaction(debtor, debtee, amountFloat, description)
	if (err != nil) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
	}
}

