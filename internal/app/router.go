package app

func (s *Server) NewRouter() {
	s.App.Get("/healthz", s.Handler.HealthCheck)

	s.App.Post("/authors", s.Handler.CreateAuthor)
	s.App.Get("/authors", s.Handler.ReadAuthors)
	s.App.Patch("/authors/:id", s.Handler.UpdateAuthor)
	s.App.Delete("/authors/:id", s.Handler.DeleteAuthor)

	s.App.Post("/books", s.Handler.CreateBook)
	s.App.Get("/books", s.Handler.ReadBooks)
	s.App.Patch("/books/:id", s.Handler.UpdateBook)
	s.App.Delete("/books/:id", s.Handler.DeleteBook)

	s.App.Post("/readers", s.Handler.CreateReader)
	s.App.Get("/readers", s.Handler.ReadReaders)
	s.App.Patch("/readers/:id", s.Handler.UpdateReader)
	s.App.Delete("/readers/:id", s.Handler.DeleteReader)

	s.App.Get("/authors/:id/books", s.Handler.GetAuthorBooks)
	s.App.Get("/readers/:id/books", s.Handler.GetReaderBooks)
}
