package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"postgrestest/models"
	"postgrestest/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_book", r.CreateBook)
	api.Put("/update_book/:id", r.UpdateBookByID)
	api.Get("/books", r.GetBooks)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Delete("/delete_book/:id", r.DeleteBookByID)
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book created successfully"})
	return nil
}

func (r *Repository) UpdateBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "id is required"})
		return nil
	}
	// Find the existing book
	bookModel := &models.Books{}
	err := r.DB.First(bookModel, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not find book"})
		return err
	}
	// Parse the request body into a new book struct
	updateData := &models.Books{}
	err = context.BodyParser(updateData)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}
	// Update fields (only the provided ones)
	err = r.DB.Model(bookModel).Updates(updateData).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not update book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book updated successfully",
		"data":    bookModel,
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Books fetched successfully",
		"data":    bookModels})
	return nil
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "id is required"})
		return nil
	}
	fmt.Println("The ID is: ", id)

	err := r.DB.Where("id=?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book fetched successfully",
		"data":    bookModel})
	return nil
}

func (r *Repository) DeleteBookByID(context *fiber.Ctx) error {
	bookModel := &models.Books{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "id is required"})
		return nil
	}
	err := r.DB.Delete(bookModel, id).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not delete book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Book deleted successfully"})
	return nil
}

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
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Could not migrate db", err)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)

	fmt.Println("Starting server on port 3000...")
	err = app.Listen(":3000")
	if err != nil {
		log.Fatal("Fiber listen failed:", err)
	}

}
