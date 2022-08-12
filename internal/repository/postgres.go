package repository

import (
	"fmt"

	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func ConnectToDB(dbHost string) (*sqlx.DB, error) {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		logrus.Fatal(err)
	}

	return db, nil
}

func CreateTableIfNotExists(db *sqlx.DB) {
	
	schema := `CREATE TABLE IF NOT EXISTS users(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null
);`
	_, err := db.Exec(schema)
	if err != nil {
		logrus.Fatal(err)
		return
	}
	logrus.Info("Created table users")
}
