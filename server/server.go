package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Server struct {
	log    zerolog.Logger
	addr   string
	router *mux.Router

	data Repository
}

func New(log zerolog.Logger, addr string, data Repository) *Server {
	srv := &Server{
		log:    log,
		addr:   addr,
		router: mux.NewRouter().StrictSlash(true),

		data: data,
	}
	srv.setupRoutes()
	return srv
}

func (s *Server) Start() error {
	httpSrv := http.Server{
		Handler:      s.router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}

	s.log.Debug().
		Str("addr", s.addr).
		Str("listen", lis.Addr().String()).
		Msg("start server")
	if err := httpSrv.Serve(lis); err != http.ErrServerClosed {
		return fmt.Errorf("serve: %w", err)
	}
	return nil
}
