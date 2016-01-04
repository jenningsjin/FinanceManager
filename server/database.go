package main

import (
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseAccessor struct {
	db *sql.DB;
}

type User struct {
	Username string
	Balance float32
	Id int
}

type Transaction struct {
	Debtor string
	Debtee string
	Amount float32
	Description string
	Timestamp string
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

	_, _, debtorId, err := dba.GetUser(debtor)
	if (err != nil) {
		return err
	}

	var debteeId int
	_, _, debteeId, err = dba.GetUser(debtee)
	if (err != nil) {
		return err
	}

	_, err = dba.db.Exec("INSERT INTO Transactions " +
		"(debtor, debtee, amount, description) VALUES " +
		"(?, ?, ?, ? )",
		debtorId, debteeId, amount, description)
	return err
}

func (dba* DatabaseAccessor) ListUsers() (map[string]float32, error) {
	ret := make(map[string]float32)
	if (dba.db == nil) {
		return ret, errors.New("Database connection is nil")
	}

	rows, queryErr := dba.db.Query("SELECT username, balance FROM Users")
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

/**
 * @brief Gets an User's information from their username
 * @param username 
 * @return string User's password
 * @return float32 User's balance
 * @return int User's id
 * @return error errors that this function has run into (nil, if no errors)
 */
func (dba* DatabaseAccessor) GetUser(username string) (string, float32, int, error) {
	if (dba.db == nil) {
		return "", 0, 0, errors.New("Database connection is nil")
	}

	row := dba.db.QueryRow("SELECT password, balance, id FROM Users WHERE username=?", username)
	var password string
	var balance float32
	var id int
	err := row.Scan(&password, &balance, &id)
	if (err != nil) {
		return "", 0, 0, err
	}
	return password, balance, id, nil
}

func (dba* DatabaseAccessor) GetTransactions() ([]Transaction, error) {
	transactions := make([]Transaction, 0)
	if (dba.db == nil) {
		return transactions, errors.New("Database connection is nil")
	}

	rows, err := dba.db.Query("SELECT Users1.username as debtorName, Users2.username as " +
		"debteeName, Transactions.amount, Transactions.description, Transactions.ts " +
		"FROM " +
		"Transactions " + 
		"LEFT JOIN " +
		"Users AS Users1 ON Transactions.debtor=Users1.id " +
		"LEFT JOIN " +
		"Users AS Users2 ON Transactions.debtee=Users2.id")
	if (err != nil) {
		return transactions, err
	}
	defer rows.Close()

	for (rows.Next()) {
		var transaction Transaction

		rowErr := rows.Scan(&transaction.Debtor, 
			&transaction.Debtee,
			&transaction.Amount,
			&transaction.Description,
			&transaction.Timestamp)
		if (rowErr != nil) {
			return transactions, rowErr
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
