package services

import (
	"api-example/src/modules/albums/models"
	"api-example/src/modules/albums/repositories"
)

type AlbumService struct {
	repository repositories.AlbumRepository
}

func (a *AlbumService) GetAlbums() ([]models.Album, error) {
	return a.repository.GetAll()
}

// func (a *AlbumService) GetAlbumByID(id int64) (models.Album, error) {
// 	return a.repository.GetAlbumByID(id)
// }
// func (a *AlbumService) CreateAlbum(album models.Album) (models.Album, error) {
// 	return a.repository.CreateAlbum(album)
// }
// func (a *AlbumService) UpdateAlbum(id int64, album models.Album) (models.Album, error) {
// 	return a.repository.UpdateAlbum(id, album)
// }
// func (a *AlbumService) DeleteAlbum(id int64) error {
// 	return a.repository.DeleteAlbum(id)
// }

func Create(repository repositories.AlbumRepository) *AlbumService {
	return &AlbumService{
		repository: repository,
	}
}