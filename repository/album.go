package repository

import (
	models "go.ut/hello/models"
	"gorm.io/gorm"
)

type AlbumRepository interface {
	FindAll(query string) []models.Album
}

type AlbumMySQLRepository struct {
	DB *gorm.DB
}

func (m *AlbumMySQLRepository) FindAll(query string) []models.Album {
	var result []models.Album
	m.DB.Raw(query).Scan(&result)
	return result
}

func NewAlbumRepository(db *gorm.DB) AlbumRepository {
	return &AlbumMySQLRepository{DB: db}
}
