package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) newHandler() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/video", s.getVideos)
	router.GET("/video/:video_id", s.getVideo)
	router.POST("/video", s.postVideo)
	router.DELETE("/video/:video_id", s.deleteVideo)
	router.PUT("/video/:video_id", s.putVideo)

	return router
}
