package encrypt

import (
	"github.com/Rolas444/apigo_base/internal/ports"
	"golang.org/x/crypto/bcrypt"
)

var _ ports.EncryptService = &EncryptionService{}

type EncryptionService struct{}

func NewEncryptionService() *EncryptionService {
	return &EncryptionService{}
}

func (s *EncryptionService) EncryptPassword(password string) (string, error) {
	hashedPAssword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPAssword), nil
}

func (s *EncryptionService) ComparePassword(hashedPassword, password string) error {

	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
