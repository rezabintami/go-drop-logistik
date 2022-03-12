package request

import (
	"go-drop-logistik/business/receipts"
	"time"
)

type Receipts struct {
	ManifestID      int       `json:"manifest_id"`
	Receiver        string    `json:"receiver"`
	PhoneReceiver   string    `json:"phone_receiver"`
	AddressReceiver string    `json:"address_receiver"`
	Sender          string    `json:"sender"`
	PhoneSender     string    `json:"phone_sender"`
	AddressSender   string    `json:"address_sender"`
	Weight          float64   `json:"weight"`
	Price           float64   `json:"price"`
	Amount          float64   `json:"amount"`
	Status          string    `json:"status"`
	PickupAt        time.Time `json:"pickup_at"`
}

func (req *Receipts) ToDomain() *receipts.Domain {
	return &receipts.Domain{
		ManifestID:      req.ManifestID,
		Receiver:        req.Receiver,
		PhoneReceiver:   req.PhoneReceiver,
		AddressReceiver: req.AddressReceiver,
		Sender:          req.Sender,
		PhoneSender:     req.PhoneSender,
		AddressSender:   req.AddressSender,
		Weight:          req.Weight,
		Price:           req.Price,
		Amount:          req.Amount,
		Status:          req.Status,
		PickupAt:        req.PickupAt,
	}
}
