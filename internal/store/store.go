package store

import (
	"Library_WebAPI/internal/model"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type Store struct {
	db     *sqlx.DB
	logger hclog.Logger
}

func NewStore(db *sqlx.DB, logger hclog.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}

func (s *Store) DatabaseCheckConnection() error {
	// Ping the database to check if the connection is alive
	if err := s.db.Ping(); err != nil {
		s.logger.Error("Database connection is down", err)
		return err
	}
	//s.logger.Debug("Database connection is up")
	return nil
}

func (s *Store) GetAuthors() ([]model.Author, error) {
	rows, err := s.db.Query(`SELECT * FROM Author`)
	if err != nil {
		s.logger.Error("Error getting authors", err)
		return nil, err
	}
	defer rows.Close()

	var authors []model.Author
	for rows.Next() {
		var author model.Author
		if err := rows.Scan(&author.Id, &author.FullName, &author.Nickname, &author.Speciality); err != nil {
			s.logger.Error("Error scanning author row", err)
			return nil, err
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		s.logger.Error("Error iterating author rows", err)
		return nil, err
	}

	return authors, nil
}

func (s *Store) CreateAuthor(author *model.Author) error {
	_, err := s.db.Exec("INSERT INTO Author (full_name, nickname, speciality) VALUES ($1, $2, $3)", author.FullName, author.Nickname, author.Speciality)
	if err != nil {
		s.logger.Error("Error creating author", err)
		return err
	}

	return nil
}

func (s *Store) UpdateAuthor(authorID int, updatedAuthor *model.Author) error {
	result, err := s.db.Exec("UPDATE Author SET full_name = $1, nickname = $2, speciality = $3 WHERE id = $4", updatedAuthor.FullName, updatedAuthor.Nickname, updatedAuthor.Speciality, authorID)
	if err != nil {
		s.logger.Error("Error updating author: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Error getting rows affected: ", err)
		return err
	}
	if rowsAffected == 0 {
		s.logger.Error("Author with ID not found: ", authorID)
		return fmt.Errorf("author with ID %d not found", authorID)
	}

	return nil
}

func (s *Store) DeleteAuthor(authorID int) error {
	result, err := s.db.Exec("DELETE FROM Author WHERE id = $1", authorID)
	if err != nil {
		s.logger.Error("Error deleting author: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Error getting rows affected: ", err)
		return err
	}
	if rowsAffected == 0 {
		s.logger.Error("Author with ID not found: ", authorID)
		return fmt.Errorf("author with ID %d not found", authorID)
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

func (s *Store) CreateBook(book *model.Book) error {
	_, err := s.db.Exec("INSERT INTO Book (title, genre, isbn, author_id) VALUES ($1, $2, $3, $4)", book.Title, book.Genre, book.ISBN, book.AuthorId)
	if err != nil {
		s.logger.Error("Error creating book", err)
		return err
	}

	return nil
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

func (s *Store) ReadReaders() ([]model.Reader, error) {
	rows, err := s.db.Query(`SELECT * FROM Reader`)
	if err != nil {
		s.logger.Error("Error getting readers", err)
		return nil, err
	}
	defer rows.Close()

	var readers []model.Reader
	for rows.Next() {
		var reader model.Reader
		var bookIDsStr string
		if err := rows.Scan(&reader.Id, &reader.FullName, &bookIDsStr); err != nil {
			s.logger.Error("Error scanning readers row", err)
			return nil, err
		}

		bookIDsStr = strings.Trim(bookIDsStr, "{}")
		bookIDs := strings.Split(bookIDsStr, ",")
		for _, idStr := range bookIDs {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				s.logger.Error("Error parsing from string to int book ID: ", err)
				return nil, err
			}
			reader.ListOfBooks = append(reader.ListOfBooks, id)
		}

		readers = append(readers, reader)
	}

	if err := rows.Err(); err != nil {
		s.logger.Error("Error iterating readers rows", err)
		return nil, err
	}

	return readers, nil
}

func (s *Store) CreateReader(reader *model.Reader) error {
	bookStr := fmt.Sprintf("{%s}", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reader.ListOfBooks)), ","), "[]"))

	_, err := s.db.Exec("INSERT INTO Reader (full_name, list_of_books) VALUES ($1, $2)", reader.FullName, bookStr)
	if err != nil {
		s.logger.Error("Error creating reader", err)
		return err
	}

	return nil
}

func (s *Store) UpdateReader(readerID int, updatedReader *model.Reader) error {
	bookStr := fmt.Sprintf("{%s}", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(updatedReader.ListOfBooks)), ","), "[]"))

	result, err := s.db.Exec("UPDATE Reader SET full_name = $1, list_of_books = $2 WHERE id = $3", updatedReader.FullName, bookStr, readerID)
	if err != nil {
		s.logger.Error("Error updating reader: ", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Error getting rows affected: ", err)
		return err
	}
	if rowsAffected == 0 {
		s.logger.Error("Reader with ID not found: ", readerID)
		return fmt.Errorf("Reader with ID %d not found", readerID)
	}

	return nil
}

func (s *Store) GetAuthorBooks(authorId int) ([]model.Book, error) {
	rows, err := s.db.Query(`SELECT * FROM Book WHERE author_id = $1`, authorId)
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
