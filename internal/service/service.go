package service

import (
	"Library_WebAPI/internal/store"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	DatabaseCheckConnection() error
}

type Service struct {
	Store  Store
	logger hclog.Logger
}

func NewService(db *sqlx.DB, log hclog.Logger) *Service {
	return &Service{
		Store:  store.NewStore(db, log),
		logger: log,
	}
}
