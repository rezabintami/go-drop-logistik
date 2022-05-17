package request

import (
	"go-drop-logistik/modules/receipts"
	"time"
)

type Receipts struct {
	ManifestID      int       `json:"manifest_id"`
	Receiver        string    `json:"receiver" validate:"required" validName:"receiver"`
	PhoneReceiver   string    `json:"phone_receiver"  validate:"required,phone,min=10,max=16" validName:"phoneNumber"`
	AddressReceiver string    `json:"address_receiver" validate:"required" validName:"address_receiver"`
	Sender          string    `json:"sender" validate:"required" validName:"sender"`
	PhoneSender     string    `json:"phone_sender"  validate:"required,phone,min=10,max=16" validName:"phoneNumber"`
	AddressSender   string    `json:"address_sender" validate:"required" validName:"address_sender"`
	Weight          float64   `json:"weight" validate:"required" validName:"weight"`
	Unit            string    `json:"unit" validate:"required" validName:"unit"`
	Price           float64   `json:"price" validate:"required" validName:"price"`
	Amount          float64   `json:"amount" validate:"required" validName:"amount"`
	Status          string    `json:"status" validate:"required" validName:"status"`
	PickupAt        time.Time `json:"pickup_at" validate:"required" validName:"pickup_at"`
}

type TrackingReceipts struct {
	Code string `json:"receipt_code"`
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
		Unit:            req.Unit,
		Price:           req.Price,
		Amount:          req.Amount,
		Status:          req.Status,
		PickupAt:        req.PickupAt,
	}
}

func (req *TrackingReceipts) TrackingReceiptToDomain() *receipts.Domain {
	return &receipts.Domain{
		Code: req.Code,
	}
}
