package superusers

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email, password string, sso bool) (string, error)
	Register(ctx context.Context, data *Domain, sso bool) error
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}
