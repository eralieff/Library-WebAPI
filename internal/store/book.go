package store

import (
	"Library_WebAPI/internal/model"
	"errors"
	"fmt"
	"strings"
)

func (s *Store) CreateBook(book *model.Book) error {
	if book.Title == "" || book.Genre == "" || book.ISBN == "" || book.AuthorId == 0 {
		s.logger.Error("Error creating book: ", "one or more empty fields")
		return errors.New("one or more empty fields")
	}

	_, err := s.db.Exec("INSERT INTO Book (title, genre, isbn, author_id) VALUES ($1, $2, $3, $4)", book.Title, book.Genre, book.ISBN, book.AuthorId)
	if err != nil {
		s.logger.Error("Error creating book", err)
		return err
	}

	return nil
}

func (s *Store) ReadBooks() ([]model.Book, error) {
	rows, err := s.db.Query(`SELECT * FROM Book`)
	if err != nil {
		s.logger.Error("Error getting books", err)
		return nil, err
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Genre, &book.ISBN, &book.AuthorId); err != nil {
			s.logger.Error("Error scanning book row", err)
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		s.logger.Error("Error iterating book rows", err)
		return nil, err
	}

	return books, nil
}

func (s *Store) UpdateBook(bookID int, updatedBook *model.Book) error {
	query := "UPDATE Book SET"
	var args []interface{}
	var paramCounter = 1

	if updatedBook.Title != "" {
		query += fmt.Sprintf(" title = $%d,", paramCounter)
		args = append(args, updatedBook.Title)
		paramCounter++
	}

	if updatedBook.Genre != "" {
		query += fmt.Sprintf(" genre = $%d,", paramCounter)
		args = append(args, updatedBook.Genre)
		paramCounter++
	}

	if updatedBook.ISBN != "" {
		query += fmt.Sprintf(" isbn = $%d,", paramCounter)
		args = append(args, updatedBook.ISBN)
		paramCounter++
	}

	if updatedBook.AuthorId != 0 {
		query += fmt.Sprintf(" author_id = $%d,", paramCounter)
		args = append(args, updatedBook.AuthorId)
		paramCounter++
	}

	if len(args) == 0 {
		s.logger.Error("Error updating book: ", "empty request")
		return errors.New("empty request")
	}

	queryString := strings.TrimSuffix(query, ",")
	queryString += fmt.Sprintf(" WHERE id = $%d", paramCounter)
	args = append(args, bookID)

	result, err := s.db.Exec(queryString, args...)
	if err != nil {
		s.logger.Error("Error updating book: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Error getting rows affected: ", err)
		return err
	}

	if rowsAffected == 0 {
		s.logger.Error("Book with ID not found: ", bookID)
		return fmt.Errorf("Book with ID %d not found", bookID)
	}

	return nil
}

func (s *Store) DeleteBook(bookID int) error {
	result, err := s.db.Exec("DELETE FROM Book WHERE id = $1", bookID)
	if err != nil {
		s.logger.Error("Error deleting book: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Error getting rows affected: ", err)
		return err
	}
	if rowsAffected == 0 {
		s.logger.Error("Book with ID not found: ", bookID)
		return fmt.Errorf("book with ID %d not found", bookID)
	}

	return nil
}
