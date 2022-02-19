package request

import (
	"go-drop-logistik/business/receipts"
)

type Receipts struct {
	Code          string  `json:"code"`
	Receiver      string  `json:"receiver"`
	Sender        string  `json:"sender"`
	PhoneReceiver string  `json:"phone_receiver"`
	PhoneSender   string  `json:"phone_sender"`
	Address       string  `json:"address"`
	Weight        float64 `json:"weight"`
	Price         float64 `json:"price"`
	Amount        float64 `json:"amount"`
	Status        string  `json:"status"`
	LicensePlate  string  `json:"license_plate"`
	DriverName    string  `json:"driver_name"`
	DriverPhone   string  `json:"driver_phone"`
}

func (req *Receipts) ToDomain() *receipts.Domain {
	return &receipts.Domain{
		Code:          req.Code,
		Receiver:      req.Receiver,
		Sender:        req.Sender,
		PhoneReceiver: req.PhoneReceiver,
		PhoneSender:   req.PhoneSender,
		Address:       req.Address,
		Weight:        req.Weight,
		Price:         req.Price,
		Amount:        req.Amount,
		Status:        req.Status,
		LicensePlate:  req.LicensePlate,
		DriverName:    req.DriverName,
		DriverPhone:   req.DriverPhone,
	}
}
