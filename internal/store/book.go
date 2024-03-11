package store

import (
	"Library_WebAPI/internal/model"
	"fmt"
)

func (s *Store) CreateBook(book *model.Book) error {
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
	result, err := s.db.Exec("UPDATE Book SET title = $1, genre = $2, isbn = $3, author_id = $4 WHERE id = $5", updatedBook.Title, updatedBook.Genre, updatedBook.ISBN, updatedBook.AuthorId, bookID)
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
		return fmt.Errorf("book with ID %d not found", bookID)
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
