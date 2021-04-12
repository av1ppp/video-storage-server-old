package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/AviParampampam/media-server/config"
)

const maxHeaderBytes = 1 << 20 // 1 MB

var isDebug = os.Getenv("DEBUG") == "true"

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
	if isDebug {
		fmt.Printf("server started on %s\n", srv.httpServer.Addr)
	}
	return srv.httpServer.ListenAndServe()
}

func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
