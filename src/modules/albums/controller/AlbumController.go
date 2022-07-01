package controller

import (
	"net/http"

	"api-example/src/modules/albums/services"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	albumService services.AlbumService
}

func (a *AlbumHandler) HandleGetAlbums(ctx *gin.Context) {
	albums, err := a.albumService.GetAlbums()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, albums)
}
