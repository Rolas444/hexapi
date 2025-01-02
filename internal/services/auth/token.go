package auth

import "github.com/Rolas444/apigo_base/internal/domain"

func (s *Service) GenerateToken(user domain.User) (string, error) {
	return "s", nil
}

func (s *Service) VerifyToken(token string) bool {
	return true
}
