package repository

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepository_ConnectToDB(t *testing.T) {

	tests := []struct {
		name          string
		dbConfig      DbConfig
		expectedError bool
	}{
		{
			name: "ok",
			dbConfig: DbConfig{
				Host:     "localhost",
				Username: "user",
				Password: "password",
				DbName:   "db",
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
				Host:     "localhost",
				Username: "user",
				Password: "password",
				DbName:   "db",
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
