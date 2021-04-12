package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (srv *Server) newHandler() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return router
}
