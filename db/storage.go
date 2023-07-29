package db

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/if3chi/PiggyVault/model"
)

type Storage interface {
	CreateAccount(*model.Account) error
	UpdateAccount(*model.Account) error
	DeleteAccount(int) error
	GetAccountByID(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

}
