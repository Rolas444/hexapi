package user

import "github.com/Rolas444/apigo_base/internal/ports"

type handler struct {
	Service ports.User
}

func NewHandler(service ports.User) *handler {
	return &handler{
		Service: service,
	}
}
