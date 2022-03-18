package drivers

import (
	"context"
	"go-drop-logistik/business/drivers"

	"gorm.io/gorm"
)

type mysqlDriverRepository struct {
	Conn *gorm.DB
}

func NewMySQLDriverRepository(conn *gorm.DB) drivers.Repository {
	return &mysqlDriverRepository{
		Conn: conn,
	}
}

func (repository *mysqlDriverRepository) Store(ctx context.Context, driverDomain *drivers.Domain) error {
	rec := fromDomain(*driverDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlDriverRepository) GetByID(ctx context.Context, id int) (drivers.Domain, error) {
	phone := Drivers{}
	result := repository.Conn.Preload("Truck").Where("id = ?", id).First(&phone)
	if result.Error != nil {
		return drivers.Domain{}, result.Error
	}

	return *phone.ToDomain(), nil
}

func (repository *mysqlDriverRepository) Update(ctx context.Context, phoneDomain *drivers.Domain, id int) error {
	phoneUpdate := fromDomain(*phoneDomain)

	result := repository.Conn.Preload("Truck").Where("id = ?", id).Updates(&phoneUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlDriverRepository) Delete(ctx context.Context, id int) error {
	phoneDelete := Drivers{}
	result := repository.Conn.Preload("Truck").Where("id = ?", id).Delete(&phoneDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
