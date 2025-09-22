package repository

import (
	"net/http"
	"postgrestest/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PublisherRepository struct {
	DB *gorm.DB
}

func (r *PublisherRepository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_publisher", r.CreatePublisher)
	api.Put("/update_publisher/:id", r.UpdatePublisherByID)
	api.Get("/publishers", r.GetPublishers)
	api.Get("/get_publisher/:id", r.GetPublisherByID)
	api.Delete("/delete_publisher/:id", r.DeletePublisherByID)
}

func (r *PublisherRepository) CreatePublisher(c *fiber.Ctx) error {
	publisher := models.Publisher{}
	if err := c.BodyParser(&publisher); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "request failed"})
	}

	if err := r.DB.Create(&publisher).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not create publisher"})
	}

	// // Preload books (initially empty)
	// if err := r.DB.Preload("Books").First(&publisher, publisher.ID).Error; err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not load publisher details"})
	// }

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Publisher created successfully",
		"data":    publisher,
	})
}

func (r *PublisherRepository) UpdatePublisherByID(c *fiber.Ctx) error {
	id := c.Params("id")
	publisher := &models.Publisher{}
	if err := r.DB.First(publisher, id).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "publisher not found"})
	}

	updateData := &models.Publisher{}
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "request failed"})
	}

	if err := r.DB.Model(publisher).Updates(updateData).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not update publisher"})
	}

	// // Preload books
	// if err := r.DB.Preload("Books").First(publisher, publisher.ID).Error; err != nil {
	// 	return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "could not load publisher details"})
	// }

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Publisher updated successfully",
		"data":    publisher,
	})
}

func (r *PublisherRepository) GetPublishers(c *fiber.Ctx) error {
	publishers := &[]models.Publisher{}
	if err := r.DB.Preload("Books").Find(publishers).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not get publishers"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Publishers fetched successfully",
		"data":    publishers,
	})
}

func (r *PublisherRepository) GetPublisherByID(c *fiber.Ctx) error {
	id := c.Params("id")
	publisher := &models.Publisher{}
	if err := r.DB.Preload("Books").First(publisher, id).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not get publisher"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Publisher fetched successfully",
		"data":    publisher,
	})
}

func (r *PublisherRepository) DeletePublisherByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := r.DB.Delete(&models.Publisher{}, id).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not delete publisher"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Publisher deleted successfully"})
}
