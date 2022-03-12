package manifestreceipt

import (
	"context"
	"go-drop-logistik/business/manifest"
	"go-drop-logistik/business/receipts"
)

type Domain struct {
	ManifestID int
	Manifest   *manifest.Domain
	ReceiptID  int
	Receipt    *receipts.Domain
}

type Usecase interface {
	Store(ctx context.Context, manifestId, ReceiptId int) error
	GetAllByManifestID(ctx context.Context, id int) ([]Domain, error)
	GetByManifestID(ctx context.Context, id int) (Domain, error)
	DeleteByReceipt(ctx context.Context, ReceiptId int) error
	DeleteByManifest(ctx context.Context, manifestId int) error
	UpdateStatusByManifest(ctx context.Context, manifestId int) error
}

type Repository interface {
	Store(ctx context.Context, manifestId, ReceiptId int) error
	GetAllByManifestID(ctx context.Context, id int) ([]Domain, error)
	GetByManifestID(ctx context.Context, id int) (Domain, error)
	DeleteByReceipt(ctx context.Context, ReceiptId int) error
	DeleteByManifest(ctx context.Context, manifestId int) error
}
