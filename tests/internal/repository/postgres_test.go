package repository

// import (
// 	"fmt"
// 	"os"
// 	"testing"

// 	"github.com/joho/godotenv"
// 	"github.com/sirupsen/logrus"
// 	"github.com/srjchsv/service/internal/repository"
// 	"github.com/stretchr/testify/require"
// )

// func init() {
// 	err := godotenv.Load("../../../.env")
// 	if err != nil {
// 		logrus.Fatal(err)
// 	}
// }
// func TestRepository_ConnectToDB(t *testing.T) {

// 	tests := []struct {
// 		name          string
// 		dbConfig      repository.DbConfig
// 		expectedError bool
// 	}{
// 		{
// 			name: "ok",
// 			dbConfig: repository.DbConfig{
// 				Host:     os.Getenv("POSTGRES_HOST"),
// 				Username: os.Getenv("POSTGRES_USER"),
// 				Password: os.Getenv("POSTGRES_PASSWORD"),
// 				DbName:   os.Getenv("POSTGRES_DB"),
// 			},
// 			expectedError: false,
// 		},
// 		{
// 			name:          "config error",
// 			dbConfig:      repository.DbConfig{},
// 			expectedError: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			db, err := repository.ConnectToDB(&test.dbConfig)
// 			defer db.Close()
// 			err = db.Ping()
// 			require.Equal(t, test.expectedError, err != nil)
// 		})
// 	}
// }

// func TestRepository_CreateTableIfNotExists(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		dbConfig      repository.DbConfig
// 		expectedError bool
// 	}{
// 		{
// 			name: "ok",
// 			dbConfig: repository.DbConfig{
// 				Host:     os.Getenv("POSTGRES_HOST"),
// 				Username: os.Getenv("POSTGRES_USER"),
// 				Password: os.Getenv("POSTGRES_PASSWORD"),
// 				DbName:   os.Getenv("POSTGRES_DB"),
// 			},
// 			expectedError: false,
// 		},
// 		{
// 			name: "config error",
// 			dbConfig: repository.DbConfig{
// 				Host:     "",
// 				Username: "user",
// 				Password: "password",
// 				DbName:   "db",
// 			},
// 			expectedError: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			db, _ := repository.ConnectToDB(&test.dbConfig)
// 			defer db.Close()
// 			err := repository.CreateTableIfNotExists(db)
// 			fmt.Println(err)
// 			require.Equal(t, test.expectedError, err != nil)

// 		})
// 	}
// }
