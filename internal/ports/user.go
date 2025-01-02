package ports

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

type User interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}
