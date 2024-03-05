package app

import (
	"Library_WebAPI/internal/model"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) HealthCheck(c *fiber.Ctx) error {
	if err := s.Service.HealthCheck(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString("Database connection is down")
	}
	return c.Status(fiber.StatusOK).SendString("OK")
}

func (s *Server) GetAuthors(c *fiber.Ctx) error {
	s.Logger.With("operation", "Get Authors")

	authors := new(model.Author)

	if err := c.BodyParser(&authors); err != nil {
		s.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON("Invalid request body")
	}

	return c.Status(fiber.StatusOK).JSON(authors)
}
