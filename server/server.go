package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/AviParampampam/video-storage-server/config"
)

const maxHeaderBytes = 1 << 20 // 1 MB

var isDebug = os.Getenv("DEBUG") == "true"

type Server struct {
	httpServer *http.Server
}

// New - создать новый сервер.
func New(conf *config.Config) *Server {
	s := &Server{}

	s.httpServer = &http.Server{
		Addr:           ":" + strconv.Itoa(conf.Server.Port),
		Handler:        s.newHandler(),
		MaxHeaderBytes: maxHeaderBytes,
		ReadTimeout:    time.Second * 10,
		WriteTimeout:   time.Second * 10,
	}

	return s
}

// Listen - начать прослушивание сервера.
func (srv *Server) Listen() error {
	if isDebug {
		fmt.Printf("server started on %s\n", srv.httpServer.Addr)
	}
	return srv.httpServer.ListenAndServe()
}

// Shutdown - завершить работу сервера
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.httpServer.Shutdown(ctx)
}
