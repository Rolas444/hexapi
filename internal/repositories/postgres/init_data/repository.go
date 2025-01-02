package initdata

import (
	"github.com/Rolas444/apigo_base/internal/ports"
	"gorm.io/gorm"
)

var _ ports.InitDataRepository = &Repository{}

type Repository struct {
	Client *gorm.DB
}

func NewRepository(client *gorm.DB) *Repository {
	return &Repository{
		Client: client,
	}
}
