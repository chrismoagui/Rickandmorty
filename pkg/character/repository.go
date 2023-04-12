package character

import "context"

// Repository handle the CRUD operations with Users.
type Repository interface {
	GetAll(ctx context.Context) ([]Character, error)
	GetOne(ctx context.Context, id uint) (Character, error)
	GetByUsername(ctx context.Context, nombre string) (Character, error)
	Create(ctx context.Context, character *Character) error
	Update(ctx context.Context, idcharacter uint, character Character) (Character, error)
	Delete(ctx context.Context, idcharacter uint) error
}
