package main

import (
	"fmt"
	"log"
	"os"
	"postgrestest/repository"
	"postgrestest/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	}

	err = storage.Migrate(db)
	if err != nil {
		log.Fatal("Could not migrate db", err)
	}

	app := fiber.New()
	// Register repositories
	bookRepo := repository.BookRepository{DB: db}
	authorRepo := repository.AuthorRepository{DB: db}
	publisherRepo := repository.PublisherRepository{DB: db}

	bookRepo.SetupRoutes(app)
	authorRepo.SetupRoutes(app)
	publisherRepo.SetupRoutes(app)

	fmt.Println("Starting server on port 3000...")
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Fiber listen failed:", err)
	}

}
