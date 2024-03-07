package app

import (
	"Library_WebAPI/internal/model"
	"github.com/gofiber/fiber/v2"
	"strconv"
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

	if len(authorsList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Author")
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

func (s *Server) UpdateAuthor(c *fiber.Ctx) error {
	s.Logger.With("operation", "Update Author")

	if c.Method() != fiber.MethodPatch {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	authorID := c.Params("id")

	author := new(model.Author)
	if err := c.BodyParser(&author); err != nil {
		s.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	id, err := strconv.Atoi(authorID)
	if err != nil {
		s.Logger.Error("Error parsing from string to int author ID: ", err)
		return err
	}

	err = s.Store.UpdateAuthor(id, author)
	if err != nil {
		//if err != ErrResourceNotFound {
		//	return c.Status(fiber.StatusNotFound).SendString("Resource Not Found")
		//}
		s.Logger.Error("Error updating author: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The author has been successfully updated")
}

func (s *Server) DeleteAuthor(c *fiber.Ctx) error {
	s.Logger.With("operation", "Delete Author")

	if c.Method() != fiber.MethodDelete {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	authorID := c.Params("id")

	id, err := strconv.Atoi(authorID)
	if err != nil {
		s.Logger.Error("Error parsing from string to int author ID: ", err)
		return err
	}

	err = s.Store.DeleteAuthor(id)
	if err != nil {
		//if err != ErrResourceNotFound {
		//	return c.Status(fiber.StatusNotFound).SendString("Resource Not Found")
		//}
		s.Logger.Error("Error deleting author: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The author has been successfully deleted")
}

func (s *Server) ReadBooks(c *fiber.Ctx) error {
	s.Logger.With("operation", "Read Books")

	if c.Method() != fiber.MethodGet {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	booksList, err := s.Store.ReadBooks()
	if err != nil {
		//if err == ErrNoBooksFound {
		//	return c.Status(fiber.StatusNotFound).SendString("No books found")
		//}
		s.Logger.Error("Error retrieving books:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	if len(booksList) == 0 {
		return c.Status(fiber.StatusOK).JSON("Empty table Books")
	}

	return c.Status(fiber.StatusOK).JSON(booksList)
}

func (s *Server) CreateBook(c *fiber.Ctx) error {
	s.Logger.With("operation", "Create Book")

	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	book := new(model.Book)
	if err := c.BodyParser(&book); err != nil {
		s.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	if err := s.Store.CreateBook(book); err != nil {
		//if err != ErrResourceNotFound {
		//	return c.Status(fiber.StatusNotFound).SendString("Resource Not Found")
		//}
		s.Logger.Error("Error creating book: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON("The book has been successfully created")
}

func (s *Server) UpdateBook(c *fiber.Ctx) error {
	s.Logger.With("operation", "Update Book")

	if c.Method() != fiber.MethodPatch {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	bookID := c.Params("id")

	book := new(model.Book)
	if err := c.BodyParser(&book); err != nil {
		s.Logger.Error("Error parsing request body: ", err)
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	id, err := strconv.Atoi(bookID)
	if err != nil {
		s.Logger.Error("Error parsing from string to int book ID: ", err)
		return err
	}

	err = s.Store.UpdateBook(id, book)
	if err != nil {
		//if err != ErrResourceNotFound {
		//	return c.Status(fiber.StatusNotFound).SendString("Resource Not Found")
		//}
		s.Logger.Error("Error updating book: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The book has been successfully updated")
}

func (s *Server) DeleteBook(c *fiber.Ctx) error {
	s.Logger.With("operation", "Delete Book")

	if c.Method() != fiber.MethodDelete {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Method Not Allowed")
	}

	bookID := c.Params("id")

	id, err := strconv.Atoi(bookID)
	if err != nil {
		s.Logger.Error("Error parsing from string to int book ID: ", err)
		return err
	}

	err = s.Store.DeleteBook(id)
	if err != nil {
		//if err != ErrResourceNotFound {
		//	return c.Status(fiber.StatusNotFound).SendString("Resource Not Found")
		//}
		s.Logger.Error("Error deleting book: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON("The book has been successfully deleted")
}
