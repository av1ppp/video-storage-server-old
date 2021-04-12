package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/AviParampampam/media-server/config"
)

const maxHeaderBytes = 1 << 20 // 1 MB

type Server struct {
	httpServer *http.Server
}

func New(conf *config.Config) *Server {
	srv := &Server{}

	srv.httpServer = &http.Server{
		Addr:           ":" + strconv.Itoa(conf.Server.Port),
		Handler:        srv.newHandler(),
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return srv
}

func (srv *Server) Listen() error {
	return srv.httpServer.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
