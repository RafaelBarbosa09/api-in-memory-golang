package main

import (
	handler "api-example/application/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handler *handler.AlbumHandler
}

func main() {
	server := &Server{
		handler: &handler.AlbumHandler{},
	}
	server.routes()
}

func (s *Server) routes() {
	router := gin.Default()
	router.GET("/albums", s.handler.HandleGetAlbums)
	router.GET("/albums/:id", s.handler.HandleGetAlbumByID)
	router.POST("/albums", s.handler.HandleSaveAlbum)
	router.PUT("/albums/:id", s.handler.HandleUpdateAlbum)
	router.DELETE("/albums/:id", s.handler.HandleDeleteAlbum)
	router.Run("localhost:8080")
}
