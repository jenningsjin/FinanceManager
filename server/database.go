package main

import (
	"errors"
	// "log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseAccessor struct {
	db *sql.DB;
}

func (dba* DatabaseAccessor) Connect(databaseUsername string, databasePassword string, databaseName string) error {
	var err error
	dba.db, err = sql.Open("mysql", databaseUsername + ":" + databasePassword + "@/" + databaseName)
	return err
}

func (dba* DatabaseAccessor) Close() {
	if (dba.db != nil) {
		dba.db.Close()
	}
}

func (dba* DatabaseAccessor) CreateUser(username string, password string) error {
	if (dba.db == nil) {
		return errors.New("Database connection is nil")
	}

	_, err := dba.db.Exec("INSERT INTO Users (username, password) VALUES (?, ?)", username, password)
	return err
}

func (dba* DatabaseAccessor) CreateTransaction(debtor string, 
	debtee string, 
	amount float32, 
	description string) error {

	if (dba.db == nil) {
		return errors.New("Database connection is nil")
	}

	_, err := dba.db.Exec("INSERT INTO Transactions " +
		"(debtor, debtee, amount, amount, description) VALUES " +
		"( (SELECT id FROM Users WHERE username=?), (SELECT id FROM Users WHERE username=?), ?, ? )",
		debtor, debtee, amount, description)
	return err
}

func (dba* DatabaseAccessor) ListUsers() (map[string]float32, error) {
	ret := make(map[string]float32)
	if (dba.db == nil) {
		return ret, errors.New("Database connection is nil")
	}

	rows, queryErr := dba.db.Query("SELECT username, balance from Users")
	if (queryErr != nil) {
		return ret, queryErr
	} 
	defer rows.Close()

	for (rows.Next()) {
		var username string
		var balance float32
		rowErr := rows.Scan(&username, &balance)
		if (rowErr != nil) {
			return ret, rowErr
		}

		ret[username] = balance
	}

	return ret, nil
}

func (dba* DatabaseAccessor) UserPassword(username string) (string, error) {
	if (dba.db == nil) {
		return "", errors.New("Database connection is nil")
	}

	row := dba.db.QueryRow("SELECT password from Users WHERE username=?", username)
	var password string
	err := row.Scan(&password)
	if (err != nil) {
		return "", err
	}
	return password, nil
}

// func (dba DatabaseAccessor) GetUser(username string) id, amount float32 {

// }

