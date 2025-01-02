package domain

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
	IdRole   uint      `json:"idRole"`
}
