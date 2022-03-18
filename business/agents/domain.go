package agents

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Password  string
	Email     string
	Phone     []string
	Roles     string
	Address   string
	Balance   float64
	Latitude  float64
	Longitude float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ExistingDomain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, email, password string, sso bool) (string, error)
	Register(ctx context.Context, data *Domain, sso bool) error
	GetByID(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, data *Domain, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
}

type Repository interface {
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByEmail(ctx context.Context, email string) (ExistingDomain, error)
	Register(ctx context.Context, data *Domain) error
	Update(ctx context.Context, data *Domain, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
}
