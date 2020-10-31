package server

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	log    zerolog.Logger
	addr   string
	router *gin.Engine

	lis     net.Listener
	httpSrv http.Server

	data Repository

	listening chan struct{}
}

func New(log zerolog.Logger, addr string, data Repository) *Server {
	srv := &Server{
		log:    log,
		addr:   addr,
		router: gin.New(),

		data: data,

		listening: make(chan struct{}),
	}
	srv.setupRoutes()
	return srv
}

func (s *Server) Start(openBrowser bool) error {
	s.httpSrv = http.Server{
		Handler:      s.router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	s.lis = lis
	if openBrowser {
		if err := browser.OpenURL("http://" + s.Addr()); err != nil {
			log.Error().
				Err(err).
				Msg("open browser")
		}
	}

	s.log.Debug().
		Str("addr", s.addr).
		Str("listen", s.Addr()).
		Msg("start server")
	close(s.listening)
	if err := s.httpSrv.Serve(lis); err != http.ErrServerClosed {
		return fmt.Errorf("serve: %w", err)
	}
	return nil
}

func (s *Server) Addr() string {
	return s.lis.Addr().String()
}

func (s *Server) Listening() <-chan struct{} {
	return s.listening
}

func (s *Server) Close() error {
	return s.httpSrv.Close()
}
