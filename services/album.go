package services

import (
	models "go.ut/hello/models"
	repositorys "go.ut/hello/repository"
)

type AlbumService interface {
	GetAll() []models.Album
}

type albumService struct {
	repo repositorys.AlbumRepository
}

func (s *albumService) GetAll() []models.Album {
	return s.repo.FindAll("select * from album")
}

func NewAlbumService(repo repositorys.AlbumRepository) AlbumService {
	return &albumService{repo: repo}
}
