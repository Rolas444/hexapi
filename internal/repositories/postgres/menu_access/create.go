package menuaccess

import (
	"context"
	"time"

	"github.com/Rolas444/apigo_base/internal/domain"
)

func (r *Repository) Create(ctx context.Context, input domain.MenuAccess) (*domain.MenuAccess, error) {
	menuAccess := &domain.MenuAccess{
		RoleId:    input.RoleId,
		MenuId:    input.Id,
		Read:      input.Read,
		Create:    input.Create,
		Update:    input.Update,
		Delete:    input.Delete,
		Status:    input.Status,
		CreatedAt: time.Now(),
	}
	result := r.Client.WithContext(ctx).Create(&menuAccess)
	if result.Error != nil {
		return nil, result.Error
	}
	return menuAccess, nil
}
