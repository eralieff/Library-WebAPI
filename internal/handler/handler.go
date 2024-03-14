package handler

import (
	"Library_WebAPI/internal/model"
	"Library_WebAPI/internal/store"
	"Library_WebAPI/pkg/validation"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	DatabaseCheckConnection() error

	CreateAuthor(author *model.Author) error
	ReadAuthors() ([]model.Author, error)
	UpdateAuthor(authorID int, updatedAuthor *model.Author) error
	DeleteAuthor(authorID int) error

	CreateBook(book *model.Book) error
	ReadBooks() ([]model.Book, error)
	UpdateBook(bookID int, updatedBook *model.Book) error
	DeleteBook(bookID int) error

	CreateReader(reader *model.Reader) error
	ReadReaders() ([]model.Reader, error)
	UpdateReader(readerID int, updatedReader *model.Reader) error
	DeleteReader(readerID int) error

	GetAuthorBooks(authorId int) ([]model.Book, error)
	GetReaderBooks(readerId int) ([]model.ReaderBook, error)
}

type IValidate interface {
	// ValidateAuthorUpdateFields(updatedAuthor *model.Author) error
}

type Handler struct {
	Store    Store
	Logger   hclog.Logger
	Validate IValidate
}

func NewHandler(db *sqlx.DB, logger hclog.Logger) *Handler {
	return &Handler{
		Store:    store.NewStore(db, logger),
		Logger:   logger,
		Validate: validation.NewValidation(logger),
	}
}
