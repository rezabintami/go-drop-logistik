package manifest

import (
	"context"
	"go-drop-logistik/business/receipts"
	"time"
)

type Domain struct {
	ID        int `gorm:"primary_key"`
	Code      string
	Status    string
	Receipt   []receipts.Domain
	DriverID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	StoreManifest(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
	Update(ctx context.Context, data *Domain, id int) error
}

type Repository interface {
	StoreManifest(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
	Update(ctx context.Context, data *Domain, id int) error
}
