package model

import (
	"math/rand"
	"time"
)

type CreateAccuntRequest struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Account struct {
	ID        int       `json:"accountId"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Number    int64     `json:"accountNumber"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        rand.Intn(10000),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}
