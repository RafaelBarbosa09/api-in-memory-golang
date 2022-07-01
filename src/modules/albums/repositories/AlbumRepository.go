package repositories

import (
	"api-example/src/modules/albums/models"
	"errors"
)

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

func (a *AlbumRepository) GetAll() (*[]models.Album, error) {
	albums := a.PopulateAlbums()

	return &albums, nil
}

func (a *AlbumRepository) GetAlbumByID(id int64) (*models.Album, error) {
	albums := a.PopulateAlbums()

	for _, album := range albums {
		if album.ID == id {
			return &album, nil
		}
	}

	return nil, errors.New("album not found")
}

func (a *AlbumRepository) CreateAlbum(album models.Album) (*models.Album, error) {
	albums := a.PopulateAlbums()

	album.ID = int64(len(albums)) + 1

	albums = append(albums, album)

	return &album, nil
}

func (a *AlbumRepository) UpdateAlbum(id int64, album models.Album) (*models.Album, error) {
	albums := a.PopulateAlbums()

	for i, album := range albums {
		if album.ID == id {
			albums[i] = album
			return &album, nil
		}
	}

	return nil, errors.New("album not found")
}

func (a *AlbumRepository) DeleteAlbum(id int64) error {
	albums := a.PopulateAlbums()

	for i, album := range albums {
		if album.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			return nil
		}
	}

	return errors.New("album not found")
}
