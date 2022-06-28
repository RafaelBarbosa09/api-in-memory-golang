package album

import "api-example/db"

type Service interface {
	GetAlbums() ([]db.Album, error)
	GetAlbumByID(id int64) (db.Album, error)
	CreateAlbum(album db.Album) (db.Album, error)
	UpdateAlbum(id int64, album db.Album) (db.Album, error)
	DeleteAlbum(id int64) error
}

type writer interface {
	CreateAlbum(album db.Album) (db.Album, error)
	UpdateAlbum(id int64, album db.Album) (db.Album, error)
	DeleteAlbum(id int64) error
}

type reader interface {
	GetAlbums() ([]db.Album, error)
	GetAlbumByID(id int64) (db.Album, error)
}

type repository interface {
	writer
	reader
}
