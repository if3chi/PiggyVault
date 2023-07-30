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
	GetAccountByID(int) (*model.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=piggy sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: db}, err
}

func (s *PostgresStore) CreateAccount(*model.Account) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(*model.Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*model.Account, error) {
	return nil, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}
