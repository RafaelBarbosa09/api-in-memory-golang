package repositories

import (
	"api-example/internal/entities"
	"errors"
)

type AlbumRepository struct {
	Albums []entities.Album
}

func (a *AlbumRepository) GetAll() ([]entities.Album, error) {
	if len(a.Albums) == 0 {
		return []entities.Album{}, nil
	}

	return a.Albums, nil
}

func (a *AlbumRepository) GetAlbumByID(id int64) (*entities.Album, error) {
	for _, album := range a.Albums {
		if album.ID == id {
			return &album, nil
		}
	}

	return nil, errors.New("album not found")
}

func (a *AlbumRepository) CreateAlbum(album entities.Album) (*entities.Album, error) {
	albums := a.Albums

	album.ID = int64(len(albums)) + 1

	a.Albums = append(albums, album)

	return &album, nil
}

func (a *AlbumRepository) UpdateAlbum(id int64, updatedAlbum entities.Album) (*entities.Album, error) {
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
