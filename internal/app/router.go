package app

func (s *Server) NewRouter() {
	s.App.Get("/healthz", s.HealthCheck)

	s.App.Get("/authors", s.GetAuthors)
	s.App.Post("/authors", s.CreateAuthor)
	s.App.Patch("/authors/:id", s.UpdateAuthor)
	s.App.Delete("/authors/:id", s.DeleteAuthor)

	s.App.Get("/books", s.ReadBooks)
	s.App.Post("/books", s.CreateBook)
	s.App.Patch("/books/:id", s.UpdateBook)
	s.App.Delete("/books/:id", s.DeleteBook)

	s.App.Get("/readers", s.ReadReaders)
	s.App.Post("/readers", s.CreateReader)
	s.App.Patch("/readers/:id", s.UpdateReader)
	s.App.Delete("/readers/:id", s.DeleteReader)

	s.App.Get("/authors/:id/books", s.GetAuthorBooks)
}
