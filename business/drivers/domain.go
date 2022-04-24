package drivers

import (
	"context"
	"go-drop-logistik/business/trucks"
	"time"
)

type Domain struct {
	ID      int
	Name    string
	Phone   string
	Address string
	TruckID int
	Truck   *trucks.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, id int) error
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, id int) error
}
