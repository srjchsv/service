package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/srjchsv/service/internal/handler"
	"github.com/srjchsv/service/internal/repository"
	"github.com/srjchsv/service/internal/services"
)

func main() {
	// Load configs
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
	dbConfig := repository.DbConfig{
		Host:     os.Getenv("POSTGRES_HOST"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
	}
	db, err := repository.ConnectToDB(&dbConfig)
	//Connect to db
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()
	//Create table if not exits
	err = repository.CreateTableIfNotExists(db)
	if err != nil {
		log.Fatal(err)
	}

	// --Initialize multi layer clean architecture structure--
	// repos is taking care of storing things
	repos := repository.NewRepository(db)
	// services  taking care of business
	services := services.NewService(repos)
	// top level handles requests & responses and routing
	handlers := handler.NewHandler(services)

	//Initialize a server instance
	r := gin.Default()
	// Initialize router
	handlers.InitRouter(r)

	//Run server
	logrus.Fatal(r.Run(os.Getenv("PORT")))
}
