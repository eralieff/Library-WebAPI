package handler

import (
	"Library_WebAPI/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) ReadReaders(c *fiber.Ctx) error {
	h.Logger.With("operation", "Read Readers")

	readersList, err := h.Store.ReadReaders()
	if err != nil {
		h.Logger.Error("Error retrieving readers:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if len(readersList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Readers")
	}

	return c.Status(fiber.StatusOK).JSON(readersList)
}

func (h *Handler) CreateReader(c *fiber.Ctx) error {
	h.Logger.With("operation", "Create Reader")

	reader := new(model.Reader)
	if err := c.BodyParser(&reader); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := h.Store.CreateReader(reader); err != nil {
		h.Logger.Error("Error creating reader: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("The reader has been successfully created")
}

func (h *Handler) UpdateReader(c *fiber.Ctx) error {
	h.Logger.With("operation", "Update Reader")

	readerID := c.Params("id")

	reader := new(model.Reader)
	if err := c.BodyParser(&reader); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	id, err := strconv.Atoi(readerID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int reader ID: ", err)
		return err
	}

	err = h.Store.UpdateReader(id, reader)
	if err != nil {
		h.Logger.Error("Error updating reader: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The reader has been successfully updated")
}

func (h *Handler) DeleteReader(c *fiber.Ctx) error {
	h.Logger.With("operation", "Delete Reader")

	readerID := c.Params("id")

	id, err := strconv.Atoi(readerID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int reader ID: ", err)
		return err
	}

	err = h.Store.DeleteReader(id)
	if err != nil {
		h.Logger.Error("Error deleting reader: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The reader has been successfully deleted")
}
