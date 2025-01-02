package user

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

func (r *Repository) Create(ctx context.Context, input domain.User) (*domain.User, error) {
	encryptedPassword, err := r.Encrypt.EncryptPassword(input.Password)
	if err != nil {
		return nil, err
	}
	user := &domain.User{
		Id:       input.Id,
		Name:     input.Name,
		UserName: input.UserName,
		Password: encryptedPassword,
	}
	result := r.Client.WithContext(ctx).Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
