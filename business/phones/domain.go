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
	StorePhone(ctx context.Context, data *Domain, id int) error
	GetByID(ctx context.Context, id int) (Domain, error)
	GetAll(ctx context.Context, agentId int) ([]Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, agentId, phoneId int) error
}

type Repository interface {
	StorePhone(ctx context.Context, data *Domain) (int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Delete(ctx context.Context, id int) error
}
