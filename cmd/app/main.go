package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
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
		Port:     os.Getenv("POSTGRES_PORT"),
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
	repository.CreateTableIfNotExists(db)

	// Initialize dependent app structure
	repos := repository.NewRepository(db)
	services := services.NewService(repos)
	handlers := handler.NewHandler(services)
	
	//Initialize a server instance
	app := fiber.New()
	// Initialize router
	handlers.InitRouter(app)

	//Run server
	logrus.Fatal(app.Listen(os.Getenv("PORT")))
}
