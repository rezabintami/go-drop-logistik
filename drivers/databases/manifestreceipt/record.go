package manifestreceipt

import (
	"go-drop-logistik/drivers/databases/manifest"
	"go-drop-logistik/drivers/databases/receipts"
	"go-drop-logistik/modules/manifestreceipt"
)

type ManifestReceipt struct {
	ManifestID int
	Manifest   *manifest.Manifest `gorm:"foreignkey:ManifestID;references:ID"`
	ReceiptID  int
	Receipt    *receipts.Receipts `gorm:"foreignkey:ReceiptID;references:ID"`
}

func fromDomain(manifestReceiptDomain manifestreceipt.Domain) *ManifestReceipt {
	return &ManifestReceipt{
		ManifestID: manifestReceiptDomain.ManifestID,
		ReceiptID:  manifestReceiptDomain.ReceiptID,
	}
}

func (rec *ManifestReceipt) ToDomain() *manifestreceipt.Domain {
	return &manifestreceipt.Domain{
		ManifestID: rec.ManifestID,
		Manifest:   rec.Manifest.ToDomain(),
		ReceiptID:  rec.ReceiptID,
		Receipt:    rec.Receipt.ToDomain(),
	}
}
