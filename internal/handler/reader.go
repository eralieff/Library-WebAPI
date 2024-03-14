package handler

import (
	"Library_WebAPI/internal/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func (h *Handler) CreateReader(c *fiber.Ctx) error {
	h.Logger.With("operation", "Create Reader")

	reader := new(model.Reader)
	if err := c.BodyParser(&reader); err != nil {
		h.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := h.Store.CreateReader(reader); err != nil {
		h.Logger.Error("Error creating reader: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).SendString("")
}

func (h *Handler) ReadReaders(c *fiber.Ctx) error {
	h.Logger.With("operation", "Read Readers")

	readersList, err := h.Store.ReadReaders()
	if err != nil {
		h.Logger.Error("Error retrieving readers:", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if len(readersList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Readers")
	}

	return c.Status(fiber.StatusOK).JSON(readersList)
}

func (h *Handler) UpdateReader(c *fiber.Ctx) error {
	h.Logger.With("operation", "Update Reader")

	requestBody := c.Body()

	if err := h.Validate.ValidateUpdateFields(requestBody, validFieldsReader); err != nil {
		h.Logger.Error("Error validating reader: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	readerID := c.Params("id")
	id, err := strconv.Atoi(readerID)
	if err != nil {
		h.Logger.Error("Error parsing from string to int reader ID: ", err)
		return err
	}

	var updatedReader model.Reader
	if err := json.Unmarshal(requestBody, &updatedReader); err != nil {
		h.Logger.Error("Error unmarshalling updated reader: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = h.Store.UpdateReader(id, &updatedReader)
	if err != nil {
		h.Logger.Error("Error updating reader: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
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
		return c.Status(fiber.StatusOK).JSON("The reader has been successfully deleted")
	}

	return c.Status(fiber.StatusOK).JSON("The reader has been successfully deleted")
}
