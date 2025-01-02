package dtos

import (
	"github.com/Rolas444/apigo_base/internal/domain"
	"github.com/google/uuid"
)

type UserDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	UserName string    `json:"username"`
	IdRole   uint      `json:"idRole"`
}

func (dto *UserDTO) FromDomain(user domain.User) {
	dto.ID = user.Id
	dto.Name = user.Name
	dto.UserName = user.UserName
	dto.IdRole = user.IdRole
}

func (dto *UserDTO) ToDomain() *domain.User {
	return &domain.User{
		Id:       dto.ID,
		Name:     dto.Name,
		UserName: dto.UserName,
		IdRole:   dto.IdRole,
	}
}

type UserCreateDTO struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	UserName string    `json:"username"`
	IdRole   uint      `json:"idRole"`
	Password string    `json:"password"`
}

func (dto *UserCreateDTO) ToDomain() *domain.User {
	return &domain.User{
		Id:       uuid.New(),
		Name:     dto.Name,
		UserName: dto.UserName,
		IdRole:   dto.IdRole,
		Password: dto.Password,
	}
}
