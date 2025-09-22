package repository

import (
	"net/http"
	"postgrestest/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	DB *gorm.DB
}

func (r *AuthorRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_author", r.CreateAuthor)
	api.Put("/update_author/:id", r.UpdateAuthorByID)
	api.Get("/authors", r.GetAuthors)
	api.Get("/get_author/:id", r.GetAuthorByID)
	api.Delete("/delete_author/:id", r.DeleteAuthorByID)
}

func (r *AuthorRepository) CreateAuthor(context *fiber.Ctx) error {
	author := models.Author{}
	if err := context.BodyParser(&author); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	if err := r.DB.Create(&author).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create author"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Author created successfully",
		"data":    author,
	})
	return nil
}

func (r *AuthorRepository) UpdateAuthorByID(context *fiber.Ctx) error {
	id := context.Params("id")
	authorModel := &models.Author{}
	if err := r.DB.First(authorModel, id).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not find author"})
		return err
	}

	updateData := &models.Author{}
	if err := context.BodyParser(updateData); err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	if err := r.DB.Model(authorModel).Updates(updateData).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not update author"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Author updated successfully",
		"data":    authorModel,
	})
	return nil
}

func (r *AuthorRepository) GetAuthors(context *fiber.Ctx) error {
	authors := &[]models.Author{}
	// Preload all books for each author
	if err := r.DB.Preload("Books").Find(authors).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get authors"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Authors fetched successfully",
		"data":    authors,
	})
	return nil
}

func (r *AuthorRepository) GetAuthorByID(context *fiber.Ctx) error {
	id := context.Params("id")
	author := &models.Author{}
	// Preload books for this author
	if err := r.DB.Preload("Books").First(author, id).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get author"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Author fetched successfully",
		"data":    author,
	})
	return nil
}

func (r *AuthorRepository) DeleteAuthorByID(context *fiber.Ctx) error {
	id := context.Params("id")
	if err := r.DB.Delete(&models.Author{}, id).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete author"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Author deleted successfully"})
	return nil
}
