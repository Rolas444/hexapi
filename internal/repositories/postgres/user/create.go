package user

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

func (r *Repository) Create(ctx context.Context, input domain.User) (*domain.User, error) {
	user := &domain.User{
		Id:       input.Id,
		Name:     input.Name,
		UserName: input.UserName,
		Password: input.Password,
	}
	result := r.Client.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
