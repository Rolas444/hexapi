package auth

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	result := r.Client.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
