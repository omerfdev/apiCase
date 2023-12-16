package main

import "math/rand"

type Account struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Status    string `json:"status"`
	Number    int64  `json:"number"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		Status:    "active",
	}
}
