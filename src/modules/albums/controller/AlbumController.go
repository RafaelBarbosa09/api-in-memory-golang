package controller

import (
	"net/http"
	"strconv"

	"api-example/src/modules/albums/models"
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

func (a *AlbumHandler) HandleGetAlbumByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := a.albumService.GetAlbumByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, album)
}

func (a *AlbumHandler) HandleSaveAlbum(ctx *gin.Context) {
	var album models.Album
	if err := ctx.BindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumCreated, err := a.albumService.CreateAlbum(album)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, albumCreated)
}

func (a *AlbumHandler) HandleUpdateAlbum(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var album models.Album
	if err := ctx.BindJSON(&album); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	albumUpdated, err := a.albumService.UpdateAlbum(id, album)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, albumUpdated)
}
