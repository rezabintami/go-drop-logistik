package receipts

import (
	"context"
	"go-drop-logistik/business/receipts"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreReceiptRepository struct {
	tx *gorm.DB
}

func NewPostgreReceiptRepository(tx *gorm.DB) receipts.Repository {
	return &postgreReceiptRepository{
		tx: tx,
	}
}

func (repository *postgreReceiptRepository) StoreReceipt(ctx context.Context, receiptDomain *receipts.Domain) (int, error) {
	rec := fromDomain(*receiptDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] receipts.repository.StoreReceipt : failed to execute store receipt query", result.Error)
		return 0, result.Error
	}
	return rec.ID, nil
}

func (repository *postgreReceiptRepository) GetByID(ctx context.Context, id int) (receipts.Domain, error) {
	rec := Receipts{}
	result := repository.tx.Where("receipts.id = ?", id).First(&rec)
	if result.Error != nil {
		log.Println("[error] receipts.repository.GetByID : failed to execute get data receipt query", result.Error)
		return receipts.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}

func (repository *postgreReceiptRepository) GetByCode(ctx context.Context, code string) (receipts.TrackDomain, error) {
	rec := Receipts{}
	result := repository.tx.Where("receipts.code = ?", code).First(&rec)
	if result.Error != nil {
		log.Println("[error] receipts.repository.GetByCode : failed to execute get data receipt query", result.Error)
		return receipts.TrackDomain{}, result.Error
	}

	return *rec.TrackToDomain(), nil
}

func (repository *postgreReceiptRepository) Fetch(ctx context.Context, page, perpage int) ([]receipts.Domain, int, error) {
	rec := []Receipts{}

	offset := (page - 1) * perpage
	err := repository.tx.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		log.Println("[error] receipts.repository.Fetch : failed to execute fetch receipts query", err)
		return []receipts.Domain{}, 0, err
	}

	var totalData int64
	err = repository.tx.Model(&rec).Count(&totalData).Error
	if err != nil {
		log.Println("[error] receipts.repository.Fetch : failed to execute count receipts query", err)
		return []receipts.Domain{}, 0, err
	}

	var result []receipts.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}

func (repository *postgreReceiptRepository) Delete(ctx context.Context, id int) error {
	rec := Receipts{}
	result := repository.tx.Where("id = ?", id).Delete(&rec)
	if result.Error != nil {
		log.Println("[error] receipts.repository.Delete : failed to execute delete receipt query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreReceiptRepository) Update(ctx context.Context, receiptDomain *receipts.Domain, id int) error {
	receiptUpdate := fromDomain(*receiptDomain)

	result := repository.tx.Where("id = ?", id).Updates(&receiptUpdate)
	if result.Error != nil {
		log.Println("[error] receipts.repository.Update : failed to execute update receipt query", result.Error)
		return result.Error
	}

	return nil
}
