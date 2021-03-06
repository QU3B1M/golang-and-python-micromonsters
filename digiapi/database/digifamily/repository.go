package digifamily

import (
	"context"
)

// Repository handle the CRUD operations with Digimon's Family.
type Repository interface {
	GetAll(ctx context.Context) []Model
	GetOne(ctx context.Context, id uint) (Model, error)
	GetByName(ctx context.Context, name string) (Model, error)
	Create(ctx context.Context, familyIn *Model) error
	Update(ctx context.Context, id uint, familyIn Model) error
	Delete(ctx context.Context, id uint) error
}
