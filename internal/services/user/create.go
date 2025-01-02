package user

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

func (s *Service) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	result, error := s.Repo.Create(ctx, *user)
	if error != nil {
		return nil, error
	}
	return result, nil
}
