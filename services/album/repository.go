package album

import "api-example/models"

type Service interface {
	GetAlbums() ([]models.Album, error)
	GetAlbumByID(id int64) (models.Album, error)
	CreateAlbum(album models.Album) (models.Album, error)
	UpdateAlbum(id int64, album models.Album) (models.Album, error)
	DeleteAlbum(id int64) error
}

type writer interface {
	CreateAlbum(album models.Album) (models.Album, error)
	UpdateAlbum(id int64, album models.Album) (models.Album, error)
	DeleteAlbum(id int64) error
}

type reader interface {
	GetAlbums() ([]models.Album, error)
	GetAlbumByID(id int64) (models.Album, error)
}

type repository interface {
	writer
	reader
}
