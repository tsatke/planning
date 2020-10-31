package server

import "github.com/gin-gonic/gin"

func (s *Server) setupRoutes() {
	rest := s.router.Group("/rest")
	rest.Use(gin.Recovery())
	rest.Use(s.Logger())
	rest.POST("/category/create", s.handleCategoryCreate())
	rest.GET("/category/delete/:category", s.handleCategoryDelete())
	rest.GET("/category/get/:category", s.handleCategoryGet())
	rest.GET("/category/rename/:categoryOld/:categoryNew", s.handleCategoryRename())
	rest.GET("/category", s.handleGetAllCategories())
}
