package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	if err := h.Store.DatabaseCheckConnection(); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString("Database connection is down")
	}
	return c.Status(fiber.StatusOK).SendString("OK")
}

func (h *Handler) GetAuthorBooks(c *fiber.Ctx) error {
	h.Logger.With("operation", "Get Author Books")

	if c.Method() != fiber.MethodGet {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	authorID := c.Params("id")

	id, err := strconv.Atoi(authorID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int author ID: ", err)
		return err
	}

	booksList, err := h.Store.GetAuthorBooks(id)
	if err != nil {
		h.Logger.Error("Error retrieving books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if len(booksList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Books")
	}

	return c.Status(fiber.StatusOK).JSON(booksList)
}

func (h *Handler) GetReaderBooks(c *fiber.Ctx) error {
	h.Logger.With("operation", "Get Reader Books")

	if c.Method() != fiber.MethodGet {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	readerID := c.Params("id")

	id, err := strconv.Atoi(readerID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int reader ID: ", err)
		return err
	}

	booksList, err := h.Store.GetReaderBooks(id)
	if err != nil {
		h.Logger.Error("Error retrieving books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if len(booksList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Books")
	}

	return c.Status(fiber.StatusOK).JSON(booksList)
}
