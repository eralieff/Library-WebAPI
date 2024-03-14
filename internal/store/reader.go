package store

import (
	"Library_WebAPI/internal/model"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func (s *Store) CreateReader(reader *model.Reader) error {
	if reader.FullName == "" || len(reader.ListOfBooks) == 0 {
		s.logger.Error("Error creating reader: ", "one or more empty fields")
		return errors.New("one or more empty fields")
	}

	bookStr := fmt.Sprintf("{%s}", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(reader.ListOfBooks)), ","), "[]"))

	_, err := s.db.Exec("INSERT INTO Reader (full_name, list_of_books) VALUES ($1, $2)", reader.FullName, bookStr)
	if err != nil {
		s.logger.Error("Error creating reader", err)
		return err
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

func (s *Store) UpdateReader(readerID int, updatedReader *model.Reader) error {
	query := "UPDATE Reader SET"
	var args []interface{}
	var paramCounter = 1

	if updatedReader.FullName != "" {
		query += fmt.Sprintf(" full_name = $%d,", paramCounter)
		args = append(args, updatedReader.FullName)
		paramCounter++
	}

	if len(updatedReader.ListOfBooks) != 0 {
		query += fmt.Sprintf(" list_of_books = $%d,", paramCounter)
		bookStr := fmt.Sprintf("{%s}", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(updatedReader.ListOfBooks)), ","), "[]"))
		args = append(args, bookStr)
		paramCounter++
	}

	if len(args) == 0 {
		s.logger.Error("Error updating reader: ", "empty request")
		return errors.New("empty request")
	}

	queryString := strings.TrimSuffix(query, ",")
	queryString += fmt.Sprintf(" WHERE id = $%d", paramCounter)
	args = append(args, readerID)

	result, err := s.db.Exec(queryString, args...)
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

func (s *Store) DeleteReader(readerID int) error {
	result, err := s.db.Exec("DELETE FROM Reader WHERE id = $1", readerID)
	if err != nil {
		s.logger.Error("Error deleting reader: ", err)
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
