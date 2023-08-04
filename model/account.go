package model

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
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
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}

func All() string {
	return `select * from account`
}

func Create(request *Account) (string, string, string, int64, int64, time.Time) {
	query := `insert into account (firstname, lastname, number, balance, created_at)
	values(?,?,?,?,?)`

	return query,
		request.FirstName, request.LastName,
		request.Number, request.Balance, request.CreatedAt
}
