package user

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

func (r *Repository) GetAll(ctx context.Context) ([]domain.User, error) {
	result := r.Client.WithContext(ctx).Limit(900).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	if result == nil {
		return nil, domain.ErrNotFound
	}

	return users, nil
}
