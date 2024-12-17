package ports

import "context"

type BaseRepository[T any, R any] interface {
	Create(ctx context.Context, input T) (*T, error)
	GetAll(ctx context.Context) ([]T, error)
	GetByID(ctx context.Context, ID R) (*T, error)
	Update(ctx context.Context, input T) error
	Delete(ctx context.Context, ID R) error
}
