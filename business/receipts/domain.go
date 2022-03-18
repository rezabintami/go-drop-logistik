package receipts

import (
	"context"
	"go-drop-logistik/business/tracks"
	"time"
)

type Domain struct {
	ID              int
	ManifestID      int
	Code            string
	Receiver        string
	PhoneReceiver   string
	AddressReceiver string
	Sender          string
	PhoneSender     string
	AddressSender   string
	Weight          float64
	Unit            string
	Price           float64
	Amount          float64
	Status          string
	PickupAt        time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TrackDomain struct {
	ID              int
	ManifestID      int
	Code            string
	Receiver        string
	PhoneReceiver   string
	AddressReceiver string
	Sender          string
	PhoneSender     string
	AddressSender   string
	Weight          float64
	Unit            string
	Price           float64
	Amount          float64
	Status          string
	Tracks          []tracks.Domain
	PickupAt        time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Usecase interface {
	StoreReceipt(ctx context.Context, data *Domain) (int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByCode(ctx context.Context, code string) (TrackDomain, error)
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
	Update(ctx context.Context, data *Domain, id int) error
}

type Repository interface {
	StoreReceipt(ctx context.Context, data *Domain) (int, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByCode(ctx context.Context, code string) (TrackDomain, error)
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, start, last int) ([]Domain, int, error)
	Update(ctx context.Context, data *Domain, id int) error
}
