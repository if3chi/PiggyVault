package model

import (
	"math/rand"
	"time"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Account struct {
	ID        int       `json:"accountId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Number    int64     `json:"accountNumber"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(firstName, lastName string) *Account {
	timeStamp := time.Now().UTC()

	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: timeStamp,
		UpdatedAt: timeStamp,
	}
}

func All() string {
	return `select * from account`
}

func Where(key, param string) (string, string, string) {
	return `select * from account where $1 = $2`, key, param
}

func Create(request *Account) (string, string, string, int64, int64, time.Time, time.Time) {
	query := `
	insert into account 
	(firstname, lastname, number, balance, created_at, updated_at) 
	values ($1, $2, $3, $4, $5, $6)`

	return query,
		request.FirstName, request.LastName,
		request.Number, request.Balance, request.CreatedAt, request.UpdatedAt
}
