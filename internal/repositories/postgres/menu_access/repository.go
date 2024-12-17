package menuaccess

import (
	"github.com/Rolas444/apigo_base/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	Client *gorm.DB
}

var menuAccess domain.MenuAccess
var listMenuAccess domain.MenuAccess

func NewRepository(client *gorm.DB) *Repository {
	return &Repository{
		Client: client,
	}
}
