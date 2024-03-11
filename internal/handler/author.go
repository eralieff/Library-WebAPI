package handler

import (
	"Library_WebAPI/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) CreateAuthor(c *fiber.Ctx) error {
	h.Logger.With("operation", "Create Author")

	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	author := new(model.Author)
	if err := c.BodyParser(&author); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := h.Store.CreateAuthor(author); err != nil {
		h.Logger.Error("Error creating author: ", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	return c.Status(fiber.StatusCreated).JSON("The author has been successfully created")
}

func (h *Handler) ReadAuthors(c *fiber.Ctx) error {
	h.Logger.With("operation", "Get Authors")

	if c.Method() != fiber.MethodGet {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	authorsList, err := h.Store.ReadAuthors()
	if err != nil {
		h.Logger.Error("Error retrieving authors:", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if len(authorsList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Author")
	}

	return c.Status(fiber.StatusOK).JSON(authorsList)
}

func (h *Handler) UpdateAuthor(c *fiber.Ctx) error {
	h.Logger.With("operation", "Update Author")

	if c.Method() != fiber.MethodPatch {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	authorID := c.Params("id")

	author := new(model.Author)
	if err := c.BodyParser(&author); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	id, err := strconv.Atoi(authorID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int author ID: ", err)
		return err
	}

	err = h.Store.UpdateAuthor(id, author)
	if err != nil {
		h.Logger.Error("Error updating author: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The author has been successfully updated")
}

func (h *Handler) DeleteAuthor(c *fiber.Ctx) error {
	h.Logger.With("operation", "Delete Author")

	if c.Method() != fiber.MethodDelete {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	authorID := c.Params("id")

	id, err := strconv.Atoi(authorID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int author ID: ", err)
		return err
	}

	err = h.Store.DeleteAuthor(id)
	if err != nil {
		h.Logger.Error("Error deleting author: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The author has been successfully deleted")
}
