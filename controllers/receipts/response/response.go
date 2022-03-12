package response

import (
	"go-drop-logistik/business/receipts"
	"time"
)

type Receipts struct {
	ID              int       `json:"id"`
	Code            string    `json:"code"`
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

type ReceiptPageResponse struct {
	Receipts *[]Receipts `json:"receipts"`
	Total    int         `json:"total"`
}

func FromDomain(receiptDomain receipts.Domain) Receipts {
	return Receipts{
		ID:              receiptDomain.ID,
		Code:            receiptDomain.Code,
		Receiver:        receiptDomain.Receiver,
		Sender:          receiptDomain.Sender,
		PhoneReceiver:   receiptDomain.PhoneReceiver,
		PhoneSender:     receiptDomain.PhoneSender,
		AddressReceiver: receiptDomain.AddressReceiver,
		AddressSender:   receiptDomain.AddressSender,
		Weight:          receiptDomain.Weight,
		Price:           receiptDomain.Price,
		Amount:          receiptDomain.Amount,
		Status:          receiptDomain.Status,
		PickupAt:        receiptDomain.PickupAt,
	}
}

func FromListDomain(receiptDomain []receipts.Domain, Count int) *ReceiptPageResponse {
	allReceipt := []Receipts{}
	for _, value := range receiptDomain {
		receipt := Receipts{
			ID:              value.ID,
			Code:            value.Code,
			Receiver:        value.Receiver,
			Sender:          value.Sender,
			PhoneReceiver:   value.PhoneReceiver,
			PhoneSender:     value.PhoneSender,
			AddressReceiver: value.AddressReceiver,
			AddressSender:   value.AddressSender,
			Weight:          value.Weight,
			Price:           value.Price,
			Amount:          value.Amount,
			Status:          value.Status,
			PickupAt:        value.PickupAt,
		}
		allReceipt = append(allReceipt, receipt)
	}
	result := ReceiptPageResponse{}
	result.Receipts = &allReceipt
	result.Total = Count
	return &result
}

func FromManifestListDomain(receiptDomain *[]receipts.Domain) (res *[]Receipts) {
	if receiptDomain != nil {
		res = &[]Receipts{}
		for _, value := range *receiptDomain {
			receipt := Receipts{
				ID:              value.ID,
				Code:            value.Code,
				Receiver:        value.Receiver,
				Sender:          value.Sender,
				PhoneReceiver:   value.PhoneReceiver,
				PhoneSender:     value.PhoneSender,
				AddressReceiver: value.AddressReceiver,
				AddressSender:   value.AddressSender,
				Weight:          value.Weight,
				Price:           value.Price,
				Amount:          value.Amount,
				Status:          value.Status,
				PickupAt:        value.PickupAt,
			}
			*res = append(*res, receipt)
		}
	}
	return res
}
