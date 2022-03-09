package receipts

import (
	"go-drop-logistik/business/receipts"
	"time"

	"gorm.io/gorm"
)

type Receipts struct {
	ID              int `gorm:"primary_key"`
	Code            string
	Receiver        string
	PhoneReceiver   string
	AddressReceiver string
	Sender          string
	PhoneSender     string
	AddressSender   string
	Weight          float64
	Price           float64
	Amount          float64
	Status          string
	PickupAt        time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (rec *Receipts) ToDomain() *receipts.Domain {
	return &receipts.Domain{
		ID:              rec.ID,
		Code:            rec.Code,
		Receiver:        rec.Receiver,
		PhoneReceiver:   rec.PhoneReceiver,
		AddressReceiver: rec.AddressReceiver,
		Sender:          rec.Sender,
		PhoneSender:     rec.PhoneSender,
		AddressSender:   rec.AddressSender,
		Weight:          rec.Weight,
		Price:           rec.Price,
		Amount:          rec.Amount,
		Status:          rec.Status,
		PickupAt:        rec.PickupAt,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func fromDomain(receiptDomain receipts.Domain) *Receipts {
	return &Receipts{
		ID:              receiptDomain.ID,
		Code:            receiptDomain.Code,
		Receiver:        receiptDomain.Receiver,
		PhoneReceiver:   receiptDomain.PhoneReceiver,
		AddressReceiver: receiptDomain.AddressReceiver,
		Sender:          receiptDomain.Sender,
		PhoneSender:     receiptDomain.PhoneSender,
		AddressSender:   receiptDomain.AddressSender,
		Weight:          receiptDomain.Weight,
		Price:           receiptDomain.Price,
		Amount:          receiptDomain.Amount,
		Status:          receiptDomain.Status,
		PickupAt:        receiptDomain.PickupAt,
		CreatedAt:       receiptDomain.CreatedAt,
		UpdatedAt:       receiptDomain.UpdatedAt,
	}
}
