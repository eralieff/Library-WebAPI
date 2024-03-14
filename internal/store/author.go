package store

import (
	"Library_WebAPI/internal/model"
	"errors"
	"fmt"
	"strings"
)

func (s *Store) CreateAuthor(author *model.Author) error {
	if author.FullName == "" || author.Nickname == "" || author.Speciality == "" {
		s.logger.Error("Error creating author: ", "one or more empty fields")
		return errors.New("one or more empty fields")
	}

	_, err := s.db.Exec("INSERT INTO Author (full_name, nickname, speciality) VALUES ($1, $2, $3)", author.FullName, author.Nickname, author.Speciality)
	if err != nil {
		s.logger.Error("Error creating author", err)
		return err
	}

	return nil
}

func (s *Store) ReadAuthors() ([]model.Author, error) {
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

// OLD: UpdateAuthor
/*
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
*/

func (s *Store) UpdateAuthor(authorID int, updatedAuthor *model.Author) error {
	query := "UPDATE Author SET"
	var args []interface{}
	var paramCounter = 1

	if updatedAuthor.FullName != "" {
		query += fmt.Sprintf(" full_name = $%d,", paramCounter)
		args = append(args, updatedAuthor.FullName)
		paramCounter++
	}

	if updatedAuthor.Nickname != "" {
		query += fmt.Sprintf(" nickname = $%d,", paramCounter)
		args = append(args, updatedAuthor.Nickname)
		paramCounter++
	}

	if updatedAuthor.Speciality != "" {
		query += fmt.Sprintf(" speciality = $%d,", paramCounter)
		args = append(args, updatedAuthor.Speciality)
		paramCounter++
	}

	if len(args) == 0 {
		s.logger.Error("Error updating author: ", "empty request")
		return errors.New("empty request")
	}

	queryString := strings.TrimSuffix(query, ",")
	queryString += fmt.Sprintf(" WHERE id = $%d", paramCounter)
	args = append(args, authorID)

	result, err := s.db.Exec(queryString, args...)
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
