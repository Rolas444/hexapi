package menu

import "gorm.io/gorm"

type Repository struct {
	Client *gorm.DB
}

func NewRepository(client *gorm.DB) *Repository {
	return &Repository{
		Client: client,
	}
}
