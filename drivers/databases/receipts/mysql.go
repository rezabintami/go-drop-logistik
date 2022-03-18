package receipts

import (
	"context"
	"go-drop-logistik/business/receipts"

	"gorm.io/gorm"
)

type mysqlReceiptRepository struct {
	Conn *gorm.DB
}

func NewMySQLReceiptRepository(conn *gorm.DB) receipts.Repository {
	return &mysqlReceiptRepository{
		Conn: conn,
	}
}

func (repository *mysqlReceiptRepository) StoreReceipt(ctx context.Context, receiptDomain *receipts.Domain) (int, error) {
	rec := fromDomain(*receiptDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return 0, result.Error
	}
	return rec.ID, nil
}

func (repository *mysqlReceiptRepository) GetByID(ctx context.Context, id int) (receipts.Domain, error) {
	rec := Receipts{}
	result := repository.Conn.Where("receipts.id = ?", id).First(&rec)
	if result.Error != nil {
		return receipts.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}

func (repository *mysqlReceiptRepository) GetByCode(ctx context.Context, code string) (receipts.TrackDomain, error) {
	rec := Receipts{}
	result := repository.Conn.Where("receipts.code = ?", code).First(&rec)
	if result.Error != nil {
		return receipts.TrackDomain{}, result.Error
	}

	return *rec.TrackToDomain(), nil
}

func (repository *mysqlReceiptRepository) Fetch(ctx context.Context, page, perpage int) ([]receipts.Domain, int, error) {
	rec := []Receipts{}

	offset := (page - 1) * perpage
	err := repository.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []receipts.Domain{}, 0, err
	}

	var totalData int64
	err = repository.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []receipts.Domain{}, 0, err
	}

	var result []receipts.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}

func (repository *mysqlReceiptRepository) Delete(ctx context.Context, id int) error {
	rec := Receipts{}
	result := repository.Conn.Where("id = ?", id).Delete(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlReceiptRepository) Update(ctx context.Context, receiptDomain *receipts.Domain, id int) error {
	receiptUpdate := fromDomain(*receiptDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&receiptUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
