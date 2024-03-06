package app

func (s *Server) NewRouter() {
	s.App.Get("/healthz", s.HealthCheck)

	s.App.Get("/authors", s.GetAuthors)
	s.App.Post("/authors", s.CreateAuthor)
	s.App.Patch("/authors/:id", s.UpdateAuthor)
}
