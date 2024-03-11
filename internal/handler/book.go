package handler

import (
	"Library_WebAPI/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) CreateBook(c *fiber.Ctx) error {
	h.Logger.With("operation", "Create Book")

	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	book := new(model.Book)
	if err := c.BodyParser(&book); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := h.Store.CreateBook(book); err != nil {
		h.Logger.Error("Error creating book: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("The book has been successfully created")
}

func (h *Handler) ReadBooks(c *fiber.Ctx) error {
	h.Logger.With("operation", "Read Books")

	if c.Method() != fiber.MethodGet {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	booksList, err := h.Store.ReadBooks()
	if err != nil {
		h.Logger.Error("Error retrieving books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if len(booksList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Books")
	}

	return c.Status(fiber.StatusOK).JSON(booksList)
}

func (h *Handler) UpdateBook(c *fiber.Ctx) error {
	h.Logger.With("operation", "Update Book")

	if c.Method() != fiber.MethodPatch {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	bookID := c.Params("id")

	book := new(model.Book)
	if err := c.BodyParser(&book); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	id, err := strconv.Atoi(bookID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int book ID: ", err)
		return err
	}

	err = h.Store.UpdateBook(id, book)
	if err != nil {
		h.Logger.Error("Error updating book: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The book has been successfully updated")
}

func (h *Handler) DeleteBook(c *fiber.Ctx) error {
	h.Logger.With("operation", "Delete Book")

	if c.Method() != fiber.MethodDelete {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	bookID := c.Params("id")

	id, err := strconv.Atoi(bookID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int book ID: ", err)
		return err
	}

	err = h.Store.DeleteBook(id)
	if err != nil {
		h.Logger.Error("Error deleting book: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The book has been successfully deleted")
}
