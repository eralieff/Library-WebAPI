package handler

import (
	"Library_WebAPI/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	h.Logger.With("operation", "Create Book")

	book := new(model.Book)
	if err := c.BodyParser(&book); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.Store.CreateBook(book); err != nil {
		h.Logger.Error("Error creating book: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).SendString("")
}

func (h *Handler) ReadBooks(c *fiber.Ctx) error {
	h.Logger.With("operation", "Read Books")

	booksList, err := h.Store.ReadBooks()
	if err != nil {
		h.Logger.Error("Error retrieving books:", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if len(booksList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Books")
	}

	return c.Status(fiber.StatusOK).JSON(booksList)
}

func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	h.Logger.With("operation", "Update Book")

	requestBody := c.Body()

	if err := h.Validate.ValidateUpdateFields(requestBody, validFieldsBook); err != nil {
		h.Logger.Error("Error validating book: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	bookID := c.Params("id")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int book ID: ", err)
		return err
	}

	var updatedBook model.Book
	if err := json.Unmarshal(requestBody, &updatedBook); err != nil {
		h.Logger.Error("Error unmarshalling updated book: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = h.Store.UpdateBook(id, &updatedBook)
	if err != nil {
		h.Logger.Error("Error updating book: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The book has been successfully updated")
}

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	h.Logger.With("operation", "Delete Book")

	bookID := c.Params("id")

	id, err := strconv.Atoi(bookID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int book ID: ", err)
		return err
	}

	err = h.Store.DeleteBook(id)
	if err != nil {
		h.Logger.Error("Error deleting book: ", err)
		return c.Status(fiber.StatusOK).JSON("The author has been successfully deleted")
	}

	return c.Status(fiber.StatusOK).JSON("The book has been successfully deleted")
}
