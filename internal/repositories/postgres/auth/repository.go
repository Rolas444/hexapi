package auth

import (
	"github.com/Rolas444/apigo_base/internal/domain"
	"github.com/Rolas444/apigo_base/internal/ports"
	"gorm.io/gorm"
)

var _ ports.AuthRepository = &Repository{}

var user domain.User

type Repository struct {
	Client *gorm.DB
}

func NewREpository(client *gorm.DB) *Repository {
	return &Repository{
		Client: client,
	}
}
