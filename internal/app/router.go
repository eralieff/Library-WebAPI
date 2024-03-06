package app

func (s *Server) NewRouter() {
	s.App.Get("/healthz", s.HealthCheck)

	s.App.Get("/authors", s.GetAuthors)
}
