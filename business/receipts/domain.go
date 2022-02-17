package receipts

import (
	"time"
)

type Domain struct {
	ID            int
	Code          string
	Receiver      string
	Sender        string
	PhoneReceiver string
	PhoneSender   string
	Address       string
	Weight        float64
	Price         float64
	Amount        float64
	Status        string
	LicensePlate  string
	DriverName    string
	DriverPhone   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Usecase interface {
}

type Repository interface {
}
