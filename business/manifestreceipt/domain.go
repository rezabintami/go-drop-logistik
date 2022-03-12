package manifestreceipt

import (
	"context"
)

type Domain struct {
	ManifestID int
	ReceiptID  int
}

type Repository interface {
	Store(ctx context.Context, phoneId, agentId int) error
	GetAllByManifestID(ctx context.Context, id int) ([]Domain, error)
	GetByManifestID(ctx context.Context, id int) (Domain, error)
	Delete(ctx context.Context, agentId, phoneId int) error
}