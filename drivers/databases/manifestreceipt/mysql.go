package manifestreceipt

import (
	"context"
	"go-drop-logistik/business/manifestreceipt"

	"gorm.io/gorm"
)

type mysqlManifestReceiptRepository struct {
	Conn *gorm.DB
}

func NewMySQLManifestReceiptRepository(conn *gorm.DB) manifestreceipt.Repository {
	return &mysqlManifestReceiptRepository{
		Conn: conn,
	}
}

func (repository *mysqlManifestReceiptRepository) Store(ctx context.Context, manifestId, ReceiptId int) error {
	manifestReceipt := &manifestreceipt.Domain{
		ManifestID: manifestId,
		ReceiptID:  ReceiptId,
	}

	rec := fromDomain(*manifestReceipt)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlManifestReceiptRepository) GetByManifestID(ctx context.Context, id int) (manifestreceipt.Domain, error) {
	manifestReceipt := ManifestReceipt{}
	result := repository.Conn.Preload("Receipt").Where("manifest_id = ?", id).First(&manifestReceipt)
	if result.Error != nil {
		return manifestreceipt.Domain{}, result.Error
	}

	return *manifestReceipt.ToDomain(), nil
}

func (repository *mysqlManifestReceiptRepository) GetByReceiptID(ctx context.Context, id int) (int, error) {
	manifestReceipt := ManifestReceipt{}
	result := repository.Conn.Where("receipt_id = ?", id).First(&manifestReceipt)
	if result.Error != nil {
		return 0, result.Error
	}

	return manifestReceipt.ManifestID, nil
}

func (repository *mysqlManifestReceiptRepository) GetAllByManifestID(ctx context.Context, id int) ([]manifestreceipt.Domain, error) {
	allManifestReceipt := []ManifestReceipt{}

	result := repository.Conn.Preload("Receipt").Where("manifest_id = ?", id).Find(&allManifestReceipt)
	if result.Error != nil {
		return []manifestreceipt.Domain{}, result.Error
	}

	allManifestReceiptDomain := []manifestreceipt.Domain{}
	for _, value := range allManifestReceipt {
		allManifestReceiptDomain = append(allManifestReceiptDomain, *value.ToDomain())
	}

	return allManifestReceiptDomain, nil
}

func (repository *mysqlManifestReceiptRepository) DeleteByReceipt(ctx context.Context, ReceiptId int) error {
	manifestReceipt := ManifestReceipt{}
	result := repository.Conn.Preload("Receipt").Preload("Manifest").Where("receipt_id = ?", ReceiptId).Delete(&manifestReceipt)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


func (repository *mysqlManifestReceiptRepository) DeleteByManifest(ctx context.Context, manifestId int) error {
	manifestReceipt := ManifestReceipt{}
	result := repository.Conn.Preload("Receipt").Preload("Manifest").Where("manifest_id = ?", manifestId).Delete(&manifestReceipt)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
