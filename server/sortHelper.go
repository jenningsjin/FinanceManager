package main

import (
	"time"
	"log"
)

type UserSlice []User

func (slice UserSlice) Len() int {
    return len(slice)
}

func (slice UserSlice) Less(i, j int) bool {
    return slice[i].Balance < slice[j].Balance
}

func (slice UserSlice) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

type TransactionByTime []Transaction

func (slice TransactionByTime) Len() int {
    return len(slice)
}

func (slice TransactionByTime) Less(i, j int) bool {
	format := "2006-01-02 15:04:05"

	timeI, err1 := time.Parse(format, slice[i].Timestamp)
	if (err1 != nil) {
		log.Printf("Time Parse Error: %s\n", err1)
	}
	timeJ, err2 := time.Parse(format, slice[j].Timestamp)
	if (err2 != nil) {
		log.Printf("Time Parse Error: %s\n", err2)
	}
	return timeI.After(timeJ)
}

func (slice TransactionByTime) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}


