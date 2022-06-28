package album

import "api-example/db"

type album struct {
	repository repository
}

func (a *album) GetAlbums() ([]db.Album, error) {
	return a.repository.GetAlbums()
}
func (a *album) GetAlbumByID(id int64) (db.Album, error) {
	return a.repository.GetAlbumByID(id)
}
func (a *album) CreateAlbum(album db.Album) (db.Album, error) {
	return a.repository.CreateAlbum(album)
}
func (a *album) UpdateAlbum(id int64, album db.Album) (db.Album, error) {
	return a.repository.UpdateAlbum(id, album)
}
func (a *album) DeleteAlbum(id int64) error {
	return a.repository.DeleteAlbum(id)
}

func Create(repository repository) Service {
	return &album{
		repository: repository,
	}
}
