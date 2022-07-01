package repositories

import "api-example/src/modules/albums/models"

type AlbumRepository struct {
	Albums []models.Album
}

func (a *AlbumRepository) PopulateAlbums() []models.Album {
	a.Albums = []models.Album{
		{ID: 1, Title: "The Dark Side of the Moon", Artist: "Pink Floyd", Price: 10.99},
		{ID: 2, Title: "The Wall", Artist: "Pink Floyd", Price: 10.99},
		{ID: 3, Title: "The Division Bell", Artist: "Pink Floyd", Price: 10.99},
	}
	return a.Albums
}

func (a *AlbumRepository) GetAll() ([]models.Album, error) {
	algums := a.PopulateAlbums()

	return algums, nil
}

// GetAlbums() ([]models.Album, error)
// GetAlbumByID(id int64) (models.Album, error)
// CreateAlbum(album models.Album) (models.Album, error)
// UpdateAlbum(id int64, album models.Album) (models.Album, error)
// DeleteAlbum(id int64) error
