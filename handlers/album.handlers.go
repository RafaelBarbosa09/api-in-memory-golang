package handlers

import (
	"net/http"

	"api-example/services/album"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	albumService album.Service
}

func (a *AlbumHandler) HandleGetAlbums(ctx *gin.Context) {
	albums, err := a.albumService.GetAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, albums)
}
