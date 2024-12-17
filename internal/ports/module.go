package ports

import "github.com/Rolas444/apigo_base/internal/domain"

type ModuleRepository interface {
	Create(module *domain.Module) (domain.Module, error)
	GetAll() ([]domain.Module, error)
}
