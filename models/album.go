package model

type Album struct {
	ID     int64 `gorm:"primary_key"`
	Title  string
	Artist string
	Price  float32
}

type Tabler interface {
	TableName() string
}

func(s *Album) TableName() string{
	return "album"
  }