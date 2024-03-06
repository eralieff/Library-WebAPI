package app

import (
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

	//authors := new(model.Author)
	//if err := c.BodyParser(&authors); err != nil {
	//	s.Logger.Error("Error parsing request body: ", err)
	//	return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	//}

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
