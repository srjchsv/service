package repository

import "github.com/jmoiron/sqlx"

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Authorization methods for repository
type Authorization interface {
	CreateUser(User) (int, error)
	GetUser(username, password string) (User, error)
}

type Repository struct {
	Authorization
}

//NewRepository creates an instance of database handler
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db)}
}
