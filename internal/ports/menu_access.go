package ports

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

type MenuAccess interface {
	Create(ctx context.Context, input *domain.MenuAccess) (domain.MenuAccess, error)
}
