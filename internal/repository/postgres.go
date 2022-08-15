package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DbConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
}

// ConnectToDB initializes db connection
func ConnectToDB(config *DbConfig) (*sqlx.DB, error) {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Username,
		config.Password,
		config.DbName)
	var err error
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateTableIfNotExist create table users if not exists
func CreateTableIfNotExists(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS users(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null
);`
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}
	return nil
}
