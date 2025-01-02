package ports

import "github.com/Rolas444/apigo_base/internal/domain"

type InitData interface {
	CreateRoles() error
	CreateDefaultUser() error
}

type InitDataRepository interface {
	CreateRole(role domain.Role) error
	CreateDefaultUser(user domain.User) error
}
