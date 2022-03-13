package trucks

import (
	"context"
	"time"
)

type Domain struct {
	ID           int
	Name         string
	Type         string
	LicensePlate string
	CreatedAt    time.Time
}

type Usecase interface {
	StoreTruck(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
}

type Repository interface {
	StoreTruck(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
}
