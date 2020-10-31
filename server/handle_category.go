package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryResponse struct {
	Name string `json:"name"`
}

func (s *Server) handleCategoryGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryName := c.Param("category")
		category := s.data.GetCategory(categoryName)
		if category != nil {
			c.JSON(http.StatusOK, response{
				Success: true,
				Element: CategoryResponse{
					Name: category.Name,
				},
			})
		} else {
			c.JSON(http.StatusNotFound, response{
				Success: false,
				Message: "category not found",
			})
		}
	}
}

func (s *Server) handleCategoryCreate() gin.HandlerFunc {
	type createRequest struct {
		Name string `json:"name"`
	}
	type createResponse struct {
		Name string `json:"name"`
	}
	return func(c *gin.Context) {
		var req createRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, response{
				Success: false,
				Message: "bad request: " + err.Error(),
			})
			return
		}

		if s.data.GetCategory(req.Name) != nil {
			c.JSON(http.StatusOK, response{
				Success: false,
				Message: "category already exists",
			})
			return
		}

		created, err := s.data.CreateCategory(req.Name)
		if err != nil {
			s.log.Error().
				Err(err).
				Msg("create category")
			c.JSON(http.StatusInternalServerError, response{
				Success: false,
				Message: "unknown error while creating",
			})
			return
		}
		c.JSON(http.StatusCreated, response{
			Success: true,
			Element: createResponse{
				Name: created.Name,
			},
		})
	}
}

func (s *Server) handleGetAllCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		categories := s.data.AllCategories()
		c.JSON(http.StatusOK, response{
			Success: len(categories) > 0,
			Element: categories,
		})
	}
}

func (s *Server) handleCategoryRename() gin.HandlerFunc {
	type renameResponse struct {
		Old CategoryResponse `json:"old"`
		New CategoryResponse `json:"new"`
	}
	return func(c *gin.Context) {
		catOld := c.Param("categoryOld")
		catNew := c.Param("categoryNew")
		if err := s.data.RenameCategory(catOld, catNew); err != nil {
			c.JSON(http.StatusBadRequest, response{
				Success: false,
				Message: err.Error(),
			})
		} else {
			category := s.data.GetCategory(catNew)
			c.JSON(http.StatusOK, response{
				Success: true,
				Element: renameResponse{
					Old: CategoryResponse{
						Name: catOld,
					},
					New: CategoryResponse{
						Name: category.Name,
					},
				},
			})
		}
	}
}

func (s *Server) handleCategoryDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		categoryName := c.Param("category")

		err := s.data.DeleteCategory(categoryName)
		var msg string
		if err != nil {
			msg = err.Error()
		}
		c.JSON(http.StatusOK, response{
			Success: err == nil,
			Message: msg,
		})
	}
}
