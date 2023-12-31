package db

import (
	"database/sql"
	"fmt"

	"github.com/if3chi/PiggyVault/model"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*model.Account) error
	UpdateAccount(*model.Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*model.Account, error)
	GetAccountByID(int) (*model.Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) Init() error {
	return s.accountTableMigration()
}

func (s *PostgresStore) accountTableMigration() error {
	query := `create table if not exists account (
		id serial primary key,
		firstname varchar(62),
		lastname varchar(62),
		number serial,
		balance serial,
		created_at timestamp,
		updated_at timestamp
	)`

	_, err := s.db.Exec(query)

	return err
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

func (s *PostgresStore) CreateAccount(account *model.Account) error {
	resp, err := s.db.Query(model.Create(account))
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) UpdateAccount(*model.Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(id int) (*model.Account, error) {
	rows, err := s.db.Query(model.Where("id", id))

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanAccount(rows)
	}

	return nil, fmt.Errorf("account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*model.Account, error) {
	rows, err := s.db.Query(model.All())

	if err != nil {
		return nil, err
	}

	accounts := []*model.Account{}

	for rows.Next() {
		account, err := scanAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)

	}

	return accounts, nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func scanAccount(row *sql.Rows) (*model.Account, error) {
	account := new(model.Account)
	err := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	return account, err
}
