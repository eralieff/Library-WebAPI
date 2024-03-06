package database

import (
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(dbStr string, logger hclog.Logger) (*sqlx.DB, error) {

	logger.Debug("create database connection")

	db, err := sqlx.Connect("postgres", dbStr)
	if err != nil {
		logger.Error("failed to create database connection", "error", err)
		return nil, err
	}

	db.SetMaxOpenConns(20)

	logger.Info("established database connection")
	return db, nil
}
