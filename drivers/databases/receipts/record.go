package receipts

import (
	"go-drop-logistik/business/receipts"
	"time"

	"gorm.io/gorm"
)

type Receipts struct {
	ID            int `gorm:"primary_key"`
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
	DeletedAt     gorm.DeletedAt
}

func (rec *Receipts) ToDomain() *receipts.Domain {
	return &receipts.Domain{
		ID:            rec.ID,
		Code:          rec.Code,
		Receiver:      rec.Receiver,
		Sender:        rec.Sender,
		PhoneReceiver: rec.PhoneReceiver,
		PhoneSender:   rec.PhoneSender,
		Address:       rec.Address,
		Weight:        rec.Weight,
		Price:         rec.Price,
		Amount:        rec.Amount,
		Status:        rec.Status,
		LicensePlate:  rec.LicensePlate,
		DriverName:    rec.DriverName,
		DriverPhone:   rec.DriverPhone,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}

func fromDomain(receiptDomain receipts.Domain) *Receipts {
	return &Receipts{
		ID:            receiptDomain.ID,
		Code:          receiptDomain.Code,
		Receiver:      receiptDomain.Receiver,
		Sender:        receiptDomain.Sender,
		PhoneReceiver: receiptDomain.PhoneReceiver,
		PhoneSender:   receiptDomain.PhoneSender,
		Address:       receiptDomain.Address,
		Weight:        receiptDomain.Weight,
		Price:         receiptDomain.Price,
		Amount:        receiptDomain.Amount,
		Status:        receiptDomain.Status,
		LicensePlate:  receiptDomain.LicensePlate,
		DriverName:    receiptDomain.DriverName,
		DriverPhone:   receiptDomain.DriverPhone,
		CreatedAt:     receiptDomain.CreatedAt,
		UpdatedAt:     receiptDomain.UpdatedAt,
	}
}
