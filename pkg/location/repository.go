package location

import "context"

// Repository handle the CRUD operations with Posts.
type Repository interface {
	GetAll(ctx context.Context) ([]Location, error)
	GetOne(ctx context.Context, id uint) (Location, error)
	GetByUser(ctx context.Context, ID uint) ([]Location, error)
	Create(ctx context.Context, location *Location) error
	Update(ctx context.Context, id uint, location Location) (Location, error)
	Delete(ctx context.Context, id uint) error
}
