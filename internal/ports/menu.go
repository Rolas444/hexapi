package ports

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

type MenuRepository interface {
	Create(ctx context.Context, menu *domain.Menu) (domain.Menu, error)
	GetByRole(idRole, idMenuAccess uint) ([]domain.Menu, error)
}
