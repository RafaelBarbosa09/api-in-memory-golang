package services

import (
	"api-example/internal/entities"
	"api-example/internal/repositories"
)

type AlbumService struct {
	repository repositories.AlbumRepository
}

func (a *AlbumService) GetAlbums() ([]entities.Album, error) {
	albums, err := a.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (a *AlbumService) GetAlbumByID(id int64) (*entities.Album, error) {
	album, err := a.repository.GetAlbumByID(id)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (a *AlbumService) CreateAlbum(album entities.Album) (*entities.Album, error) {
	newAlbum, err := a.repository.CreateAlbum(album)
	if err != nil {
		return nil, err
	}

	return newAlbum, nil
}

func (a *AlbumService) UpdateAlbum(id int64, album entities.Album) (*entities.Album, error) {
	updatedAlbum, err := a.repository.UpdateAlbum(id, album)
	if err != nil {
		return nil, err
	}

	return updatedAlbum, nil
}

func (a *AlbumService) DeleteAlbum(id int64) error {
	err := a.repository.DeleteAlbum(id)
	if err != nil {
		return err
	}

	return nil
}
