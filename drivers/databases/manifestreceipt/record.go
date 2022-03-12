package manifestreceipt

import "go-drop-logistik/business/manifestreceipt"

type ManifestReceipt struct {
	ManifestID int
	ReceiptID  int
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
		ReceiptID:  rec.ReceiptID,
	}
}
