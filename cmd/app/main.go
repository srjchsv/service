package main

import (
	"flag"
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
	dbHost := flag.String("host", "localhost", "postgres database host")
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal(err)
	}
	//Initialize a server instance
	app := fiber.New()

	//Connect to db
	db, err := repository.ConnectToDB(*dbHost)
	if err != nil {
		logrus.Fatal(err)
	}
	defer db.Close()

	//Create table if not exits
	repository.CreateTableIfNotExists(db)

	// Initialize app structure
	repos := repository.NewRepository(db)
	services := services.NewService(repos)
	handlers := handler.NewHandler(services)

	// Initialize router
	handlers.InitRouter(app)

	//Run server
	logrus.Fatal(app.Listen(os.Getenv("PORT")))
}
