package role

import (
	"github.com/Rolas444/apigo_base/internal/domain"
	"github.com/Rolas444/apigo_base/internal/ports"
	"gorm.io/gorm"
)

var _ ports.BaseRepository[domain.Role, uint] = &Repository{}

var role domain.Role
var roles []domain.Role

type Repository struct {
	Client *gorm.DB
}

func NewRepository(client *gorm.DB) *Repository {
	return &Repository{
		Client: client,
	}
}
