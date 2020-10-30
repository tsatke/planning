package server

func (s *Server) setupRoutes() {
	rest := s.router.PathPrefix("/rest").Subrouter()
	rest.Use(s.middleware)
	rest.Handle("/category/{categoryOld}/rename/{categoryNew}", s.handleCategoryRename())
	rest.Handle("/category/{category}/create", s.handleCategoryCreate())
	rest.Handle("/category/{category}/delete", s.handleCategoryDelete())
	rest.Handle("/category/{category}", s.handleCategoryGet())
	rest.Handle("/category", s.handleGetAllCategories())
}
