package repositories

import (
	"api-example/src/modules/albums/models"
	"errors"
)

type AlbumRepository struct {
	Albums []models.Album
}

func (a *AlbumRepository) GetAll() (*[]models.Album, error) {
	if len(a.Albums) == 0 {
		return &[]models.Album{}, nil
	}

	return &a.Albums, nil
}

func (a *AlbumRepository) GetAlbumByID(id int64) (*models.Album, error) {
	for _, album := range a.Albums {
		if album.ID == id {
			return &album, nil
		}
	}

	return nil, errors.New("album not found")
}

func (a *AlbumRepository) CreateAlbum(album models.Album) (*models.Album, error) {
	albums := a.Albums

	album.ID = int64(len(albums)) + 1

	a.Albums = append(albums, album)

	return &album, nil
}

func (a *AlbumRepository) UpdateAlbum(id int64, updatedAlbum models.Album) (*models.Album, error) {
	for i, album := range a.Albums {
		if album.ID == id {
			updatedAlbum.ID = album.ID
			a.Albums[i] = updatedAlbum
			return &updatedAlbum, nil
		}
	}

	return nil, errors.New("album not found")
}

func (a *AlbumRepository) DeleteAlbum(id int64) error {
	for i, album := range a.Albums {
		if album.ID == id {
			a.Albums = append(a.Albums[:i], a.Albums[i+1:]...)
			return nil
		}
	}

	return errors.New("album not found")
}
