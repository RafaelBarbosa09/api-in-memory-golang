package interfaces

import "api-example/core/models"

type AlbumPort interface {
	GetAll() (*[]models.Album, error)
	GetAlbumByID(id int64) (*models.Album, error)
	CreateAlbum(album models.Album) (*models.Album, error)
	UpdateAlbum(id int64, updatedAlbum models.Album) (*models.Album, error)
	DeleteAlbum(id int64) error
}
