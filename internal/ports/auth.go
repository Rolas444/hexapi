package ports

import (
	"context"

	"github.com/Rolas444/apigo_base/internal/domain"
)

type Auth interface {
	Login(username, password string) (string, error)
	GenerateToken(user domain.User) (string, error)
	VerifyToken(token string) bool
}

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
}
