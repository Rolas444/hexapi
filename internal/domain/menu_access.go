package domain

import (
	"time"

	"github.com/google/uuid"
)

type MenuAccess struct {
	Id        uint      `json:"id"`
	MenuId    uint      `json:"menuId"`
	RoleId    uuid.UUID `json:"roleId"`
	Read      bool
	Create    bool
	Update    bool
	Delete    bool
	Status    int        `json:"status"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
