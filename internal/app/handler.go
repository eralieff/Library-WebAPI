package app

import (
	"Library_WebAPI/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) HealthCheck(c *fiber.Ctx) error {
	if err := s.Store.DatabaseCheckConnection(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString("Database connection is down")
	}
	return c.Status(fiber.StatusOK).SendString("OK")
}

func (s *Server) GetAuthors(c *fiber.Ctx) error {
	s.Logger.With("operation", "Get Authors")

	if c.Method() != fiber.MethodGet {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	authorsList, err := s.Store.GetAuthors()
	if err != nil {
		//if err == ErrNoAuthorsFound {
		//	return c.Status(fiber.StatusNotFound).SendString("No authors found")
		//}
		s.Logger.Error("Error retrieving authors:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusOK).JSON(authorsList)
}

func (s *Server) CreateAuthor(c *fiber.Ctx) error {
	s.Logger.With("operation", "Create Author")

	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	author := new(model.Author)
	if err := c.BodyParser(&author); err != nil {
		s.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := s.Store.CreateAuthor(author); err != nil {
		//if err != ErrResourceNotFound {
		//	return c.Status(fiber.StatusNotFound).SendString("Resource Not Found")
		//}
		s.Logger.Error("Error creating author: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON("The author has been successfully created")
}
