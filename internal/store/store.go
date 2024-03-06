package store

import (
	"Library_WebAPI/internal/model"
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
