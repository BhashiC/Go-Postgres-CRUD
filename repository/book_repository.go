package repository

import (
	"net/http"
	"postgrestest/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_book", r.CreateBook)
	api.Put("/update_book/:id", r.UpdateBookByID)
	api.Get("/books", r.GetBooks)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Delete("/delete_book/:id", r.DeleteBookByID)
}

// CreateBook validates author/publisher before inserting
func (r *BookRepository) CreateBook(c *fiber.Ctx) error {
	book := models.Book{}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "invalid request"})
	}

	// Check if author exists
	var author models.Author
	if err := r.DB.First(&author, book.AuthorID).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid author ID"})
	}

	// Check if publisher exists
	var publisher models.Publisher
	if err := r.DB.First(&publisher, book.PublisherID).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid publisher ID"})
	}

	if err := r.DB.Create(&book).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not create book"})
	}

	// // Preload author and publisher for response
	// if err := r.DB.Preload("Author").Preload("Publisher").First(&book, book.ID).Error; err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not load book details"})
	// }

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book created successfully",
		"data":    book,
	})
}

func (r *BookRepository) UpdateBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "id is required"})
	}

	book := &models.Book{}
	if err := r.DB.First(book, id).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "book not found"})
	}

	updateData := &models.Book{}
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "invalid request"})
	}

	// Optional: validate author/publisher IDs if provided
	if updateData.AuthorID != 0 {
		var author models.Author
		if err := r.DB.First(&author, updateData.AuthorID).Error; err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid author ID"})
		}
	}
	if updateData.PublisherID != 0 {
		var publisher models.Publisher
		if err := r.DB.First(&publisher, updateData.PublisherID).Error; err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "invalid publisher ID"})
		}
	}

	if err := r.DB.Model(book).Updates(updateData).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not update book"})
	}

	// // Preload author and publisher for response
	// if err := r.DB.Preload("Author").Preload("Publisher").First(book, book.ID).Error; err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not load book details"})
	// }

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book updated successfully",
		"data":    book,
	})
}

func (r *BookRepository) GetBooks(c *fiber.Ctx) error {
	books := &[]models.Book{}
	if err := r.DB.Preload("Author").Preload("Publisher").Find(books).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not get books"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Books fetched successfully",
		"data":    books,
	})
}

func (r *BookRepository) GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	book := &models.Book{}
	if err := r.DB.Preload("Author").Preload("Publisher").First(book, id).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not get book"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book fetched successfully",
		"data":    book,
	})
}

func (r *BookRepository) DeleteBookByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "id is required"})
	}
	book := &models.Book{}
	if err := r.DB.Delete(book, id).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not delete book"})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Book deleted successfully"})
}
