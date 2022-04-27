package manifestreceipt

import (
	"context"
	"go-drop-logistik/business/manifestreceipt"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreManifestReceiptRepository struct {
	tx *gorm.DB
}

func NewPostgreManifestReceiptRepository(tx *gorm.DB) manifestreceipt.Repository {
	return &postgreManifestReceiptRepository{
		tx: tx,
	}
}

func (repository *postgreManifestReceiptRepository) Store(ctx context.Context, manifestId, ReceiptId int) error {
	manifestReceipt := &manifestreceipt.Domain{
		ManifestID: manifestId,
		ReceiptID:  ReceiptId,
	}

	rec := fromDomain(*manifestReceipt)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] manifestreceipt.repository.Store : failed to execute store manifestreceipt query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreManifestReceiptRepository) GetByManifestID(ctx context.Context, id int) (manifestreceipt.Domain, error) {
	manifestReceipt := ManifestReceipt{}
	result := repository.tx.Preload("Receipt").Where("manifest_id = ?", id).First(&manifestReceipt)
	if result.Error != nil {
		log.Println("[error] manifestreceipt.repository.GetByManifestID : failed to execute get data manifestreceipt query", result.Error)
		return manifestreceipt.Domain{}, result.Error
	}

	return *manifestReceipt.ToDomain(), nil
}

func (repository *postgreManifestReceiptRepository) GetByReceiptID(ctx context.Context, id int) (int, error) {
	manifestReceipt := ManifestReceipt{}
	result := repository.tx.Where("receipt_id = ?", id).First(&manifestReceipt)
	if result.Error != nil {
		log.Println("[error] manifestreceipt.repository.GetByReceiptID : failed to execute get data manifestreceipt query", result.Error)
		return 0, result.Error
	}

	return manifestReceipt.ManifestID, nil
}

func (repository *postgreManifestReceiptRepository) GetAllByManifestID(ctx context.Context, id int) ([]manifestreceipt.Domain, error) {
	allManifestReceipt := []ManifestReceipt{}

	result := repository.tx.Preload("Receipt").Where("manifest_id = ?", id).Find(&allManifestReceipt)
	if result.Error != nil {
		log.Println("[error] manifestreceipt.repository.GetAllByManifestID : failed to execute get data manifestreceipts query", result.Error)
		return []manifestreceipt.Domain{}, result.Error
	}

	allManifestReceiptDomain := []manifestreceipt.Domain{}
	for _, value := range allManifestReceipt {
		allManifestReceiptDomain = append(allManifestReceiptDomain, *value.ToDomain())
	}

	return allManifestReceiptDomain, nil
}

func (repository *postgreManifestReceiptRepository) DeleteByReceipt(ctx context.Context, ReceiptId int) error {
	manifestReceipt := ManifestReceipt{}
	result := repository.tx.Preload("Receipt").Preload("Manifest").Where("receipt_id = ?", ReceiptId).Delete(&manifestReceipt)
	if result.Error != nil {
		log.Println("[error] manifestreceipt.repository.DeleteByReceipt : failed to execute delete manifestreceipt query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreManifestReceiptRepository) DeleteByManifest(ctx context.Context, manifestId int) error {
	manifestReceipt := ManifestReceipt{}
	result := repository.tx.Preload("Receipt").Preload("Manifest").Where("manifest_id = ?", manifestId).Delete(&manifestReceipt)
	if result.Error != nil {
		log.Println("[error] manifestreceipt.repository.DeleteByManifest : failed to execute delete manifestreceipt query", result.Error)
		return result.Error
	}

	return nil
}
