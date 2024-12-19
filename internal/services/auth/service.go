package auth

import (
	"github.com/Rolas444/apigo_base/internal/domain"
	"github.com/Rolas444/apigo_base/internal/ports"
	"github.com/google/uuid"
)

var _ ports.Auth = &Service{}

type Service struct {
	Repo ports.BaseRepository[domain.User, uuid.UUID]
}

func NewService(repo ports.BaseRepository[domain.User, uuid.UUID]) *Service {
	return &Service{
		Repo: repo,
	}
}
