package phones

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	StorePhone(ctx context.Context, data *Domain, id int) (int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, phoneId int) error
}

type Repository interface {
	StorePhone(ctx context.Context, data *Domain) (int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, id int) error
}
