package drivers

import (
	"context"
	"go-drop-logistik/modules/drivers"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreDriverRepository struct {
	tx *gorm.DB
}

func NewPostgreDriverRepository(tx *gorm.DB) drivers.Repository {
	return &postgreDriverRepository{
		tx: tx,
	}
}

func (repository *postgreDriverRepository) Store(ctx context.Context, driverDomain *drivers.Domain) error {
	rec := fromDomain(*driverDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] drivers.repository.Store : failed to execute store driver query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreDriverRepository) GetByID(ctx context.Context, id int) (drivers.Domain, error) {
	phone := Drivers{}
	result := repository.tx.Preload("Truck").Where("id = ?", id).First(&phone)
	if result.Error != nil {
		log.Println("[error] drivers.repository.GetByID : failed to execute get data driver query", result.Error)
		return drivers.Domain{}, result.Error
	}

	return *phone.ToDomain(), nil
}

func (repository *postgreDriverRepository) Update(ctx context.Context, phoneDomain *drivers.Domain, id int) error {
	phoneUpdate := fromDomain(*phoneDomain)

	result := repository.tx.Preload("Truck").Where("id = ?", id).Updates(&phoneUpdate)
	if result.Error != nil {
		log.Println("[error] drivers.repository.Update : failed to execute update driver query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreDriverRepository) Delete(ctx context.Context, id int) error {
	phoneDelete := Drivers{}
	result := repository.tx.Preload("Truck").Where("id = ?", id).Delete(&phoneDelete)
	if result.Error != nil {
		log.Println("[error] drivers.repository.Delete : failed to execute delete driver query", result.Error)
		return result.Error
	}

	return nil
}
