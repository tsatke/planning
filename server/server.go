package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/browser"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

func (s *Server) Start(openBrowser bool) error {
	httpSrv := http.Server{
		Handler:      s.router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	listenerAddress := lis.Addr().String()

	if openBrowser {
		if err := browser.OpenURL("http://" + listenerAddress); err != nil {
			log.Error().
				Err(err).
				Msg("open browser")
		}
	}

	s.log.Debug().
		Str("addr", s.addr).
		Str("listen", listenerAddress).
		Msg("start server")
	if err := httpSrv.Serve(lis); err != http.ErrServerClosed {
		return fmt.Errorf("serve: %w", err)
	}
	return nil
}
