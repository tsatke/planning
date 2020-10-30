package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) handleCategoryGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		categoryName := vars["category"]
		category := s.data.GetCategory(categoryName)
		s.writeResponse(w, SingleElementRequestResponse{
			Success: category != nil,
			Element: category,
		})
	}
}

func (s *Server) handleCategoryCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		categoryName := vars["category"]
		category, err := s.data.CreateCategory(categoryName)
		if err != nil {
			s.writeResponse(w, SingleElementRequestResponse{
				Success: false,
				Message: err.Error(),
			})
		} else {
			s.writeResponse(w, SingleElementRequestResponse{
				Success: true,
				Element: category,
			})
		}
	}
}

func (s *Server) handleGetAllCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		categories := s.data.AllCategories()
		s.writeResponse(w, ManyElementsRequestResponse{
			Success:  len(categories) > 0,
			Elements: categories,
		})
	}
}

func (s *Server) handleCategoryRename() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		old := vars["categoryOld"]
		new := vars["categoryNew"]
		if err := s.data.RenameCategory(old, new); err != nil {
			s.writeResponse(w, SingleElementRequestResponse{
				Success: false,
				Message: err.Error(),
			})
		} else {
			s.writeResponse(w, SingleElementRequestResponse{
				Success: true,
				Element: s.data.GetCategory(new),
			})
		}
	}
}

func (s *Server) handleCategoryDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		categoryName := vars["category"]

		err := s.data.DeleteCategory(categoryName)
		var msg string
		if err != nil {
			msg = err.Error()
		}
		s.writeResponse(w, SingleElementRequestResponse{
			Success: err == nil,
			Message: msg,
		})
	}
}

type SingleElementRequestResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Element interface{} `json:"element,omitempty"`
}

type ManyElementsRequestResponse struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message,omitempty"`
	Elements interface{} `json:"elements,omitempty"`
}

func (s *Server) writeResponse(w http.ResponseWriter, response interface{}) {
	data, err := json.Marshal(response)
	if err != nil {
		s.log.Error().
			Err(err).
			Msg("marshal")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(data); err != nil {
		s.log.Error().
			Err(err).
			Msg("write")
		return
	}
}
