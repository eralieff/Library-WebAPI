package store

import (
	"Library_WebAPI/internal/model"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
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
