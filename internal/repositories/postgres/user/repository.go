package user

import (
	"github.com/Rolas444/apigo_base/internal/domain"
	"github.com/Rolas444/apigo_base/internal/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ ports.BaseRepository[domain.User, uuid.UUID] = &Repository{}

var user domain.User
var users []domain.User

type Repository struct {
	Client *gorm.DB
}

func NewRepository(client *gorm.DB) *Repository {
	return &Repository{
		Client: client,
	}
}
