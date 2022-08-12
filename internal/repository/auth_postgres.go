package repository

import (
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

//NewAuthPostgres return a db instance
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

//CreateUser creates user in db
func (r *AuthPostgres) CreateUser(user User) (int, error) {
	var id int

	row := r.db.QueryRow("INSERT INTO users (name,username, password) values ($1, $2, $3) RETURNING id", user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

//GetUser gets user by username and password
func (r *AuthPostgres) GetUser(username, password string) (User, error) {
	var user User
	err := r.db.Get(&user, "SELECT id FROM users WHERE username=$1 AND password=$2", username, password)

	return user, err
}
