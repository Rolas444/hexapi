package user

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
	"github.com/google/uuid"
)

func (r *Repository) GetByID(ctx context.Context, ID uuid.UUID) (*domain.User, error) {
	return nil, nil
}
