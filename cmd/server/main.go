package main

import (
	"api-example/internal/controller"

	"github.com/gin-gonic/gin"
)

type Server struct {
	controller *controller.AlbumHandler
}

func main() {
	server := &Server{
		controller: &controller.AlbumHandler{},
	}
	server.routes()
}

func (s *Server) routes() {
	router := gin.Default()
	router.GET("/albums", s.controller.HandleGetAlbums)
	router.GET("/albums/:id", s.controller.HandleGetAlbumByID)
	router.POST("/albums", s.controller.HandleSaveAlbum)
	router.PUT("/albums/:id", s.controller.HandleUpdateAlbum)
	router.DELETE("/albums/:id", s.controller.HandleDeleteAlbum)
	router.Run("localhost:8080")
}
