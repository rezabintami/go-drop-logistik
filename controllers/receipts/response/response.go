package response

import (
	"go-drop-logistik/business/receipts"
)

type Receipts struct {
	ID            int     `json:"id"`
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

type ReceiptPageResponse struct {
	Receipts *[]Receipts `json:"receipts"`
	Total    int         `json:"total"`
}

func FromDomain(receiptDomain receipts.Domain) Receipts {
	return Receipts{
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
	}
}

func FromListDomain(agentDomain []receipts.Domain, Count int) *ReceiptPageResponse {
	allReceipt := []Receipts{}
	for _, value := range agentDomain {
		agent := Receipts{
			ID:            value.ID,
			Code:          value.Code,
			Receiver:      value.Receiver,
			Sender:        value.Sender,
			PhoneReceiver: value.PhoneReceiver,
			PhoneSender:   value.PhoneSender,
			Address:       value.Address,
			Weight:        value.Weight,
			Price:         value.Price,
			Amount:        value.Amount,
			Status:        value.Status,
			LicensePlate:  value.LicensePlate,
			DriverName:    value.DriverName,
			DriverPhone:   value.DriverPhone,
		}
		allReceipt = append(allReceipt, agent)
	}
	result := ReceiptPageResponse{}
	result.Receipts = &allReceipt
	result.Total = Count
	return &result
}
