package repository

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		logrus.Fatal(err)
	}
}
func TestRepository_ConnectToDB(t *testing.T) {

	tests := []struct {
		name          string
		dbConfig      DbConfig
		expectedError bool
	}{
		{
			name: "ok",
			dbConfig: DbConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Username: os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				DbName:   os.Getenv("POSTGRES_DB"),
			},
			expectedError: false,
		},
		{
			name:          "config error",
			dbConfig:      DbConfig{},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var errBool bool
			db, err := ConnectToDB(&test.dbConfig)
			if err != nil {
				errBool = true
			}
			defer db.Close()
			err = db.Ping()
			if err != nil {
				errBool = true
			}
			require.Equal(t, test.expectedError, errBool)

		})
	}

}

func TestRepository_CreateTableIfNotExists(t *testing.T) {
	tests := []struct {
		name          string
		dbConfig      DbConfig
		expectedError bool
	}{
		{
			name: "ok",
			dbConfig: DbConfig{
				Host:     os.Getenv("POSTGRES_HOST"),
				Username: os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				DbName:   os.Getenv("POSTGRES_DB"),
			},
			expectedError: false,
		},
		{
			name: "config error",
			dbConfig: DbConfig{
				Host:     "",
				Username: "user",
				Password: "password",
				DbName:   "db",
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var errBool bool
			db, _ := ConnectToDB(&test.dbConfig)
			defer db.Close()
			err := CreateTableIfNotExists(db)
			if err != nil {
				errBool = true
			}
			require.Equal(t, test.expectedError, errBool)

		})
	}

}
