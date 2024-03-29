package services

import (
	"api-example/core/models"
	"api-example/datasources/database"
)

type AlbumService struct {
	repository database.AlbumRepository
}

func (a *AlbumService) GetAlbums() ([]models.Album, error) {
	albums, err := a.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return *albums, nil
}

func (a *AlbumService) GetAlbumByID(id int64) (*models.Album, error) {
	album, err := a.repository.GetAlbumByID(id)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (a *AlbumService) CreateAlbum(album models.Album) (*models.Album, error) {
	newAlbum, err := a.repository.CreateAlbum(album)
	if err != nil {
		return nil, err
	}

	return newAlbum, nil
}

func (a *AlbumService) UpdateAlbum(id int64, album models.Album) (*models.Album, error) {
	updatedAlbum, err := a.repository.UpdateAlbum(id, album)
	if err != nil {
		return nil, err
	}

	return updatedAlbum, nil
}

func (a *AlbumService) DeleteAlbum(id int64) error {
	if err := a.repository.DeleteAlbum(id); err != nil {
		return err
	}

	return nil
}
