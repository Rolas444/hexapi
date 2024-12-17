package role

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

func (r *Repository) Create(ctx context.Context, input domain.Role) (*domain.Role, error) {
	role := &domain.Role{
		Id:     input.Id,
		Name:   input.Name,
		Desc:   input.Desc,
		Status: input.Status,
	}
	result := r.Client.WithContext(ctx).Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil
}
