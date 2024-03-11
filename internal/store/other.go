package store

import "Library_WebAPI/internal/model"

func (s *Store) DatabaseCheckConnection() error {
	if err := s.db.Ping(); err != nil {
		s.logger.Error("Database connection is down", err)
		return err
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

func (s *Store) GetReaderBooks(readerId int) ([]model.ReaderBook, error) {
	rows, err := s.db.Query(`SELECT r.full_name AS reader_name, b.title AS book_title, b.genre, b.isbn FROM Reader r JOIN unnest(r.list_of_books) AS book_id ON true JOIN Book b ON book_id = b.id WHERE r.id = $1`, readerId)
	if err != nil {
		s.logger.Error("Error getting reader books", err)
		return nil, err
	}
	defer rows.Close()

	var readerBooks []model.ReaderBook
	for rows.Next() {
		var readerBook model.ReaderBook
		if err := rows.Scan(&readerBook.ReaderName, &readerBook.BookTitle, &readerBook.Genre, &readerBook.ISBN); err != nil {
			s.logger.Error("Error scanning reader book row", err)
			return nil, err
		}
		readerBooks = append(readerBooks, readerBook)
	}

	if err := rows.Err(); err != nil {
		s.logger.Error("Error iterating reader book rows", err)
		return nil, err
	}

	return readerBooks, nil
}
