package trucks

import (
	"context"
	"go-drop-logistik/business/trucks"

	"github.com/jinzhu/gorm"
)

type mysqlTruckRepository struct {
	Conn *gorm.DB
}

func NewMySQLTruckRepository(conn *gorm.DB) trucks.Repository {
	return &mysqlTruckRepository{
		Conn: conn,
	}
}

func (repository *mysqlTruckRepository) StoreTruck(ctx context.Context, truckDomain *trucks.Domain) error {
	rec := fromDomain(*truckDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTruckRepository) GetByID(ctx context.Context, id int) (trucks.Domain, error) {
	truck := Trucks{}
	result := repository.Conn.Where("id = ?", id).First(&truck)
	if result.Error != nil {
		return trucks.Domain{}, result.Error
	}

	return *truck.ToDomain(), nil
}

func (repository *mysqlTruckRepository) Update(ctx context.Context, truckDomain *trucks.Domain, id int) error {
	truckUpdate := fromDomain(*truckDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&truckUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTruckRepository) Delete(ctx context.Context, id int) error {
	truckDelete := Trucks{}
	result := repository.Conn.Where("id = ?", id).Delete(&truckDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTruckRepository) Fetch(ctx context.Context, page, perpage int) ([]trucks.Domain, int, error) {
	rec := []Trucks{}

	offset := (page - 1) * perpage
	err := repository.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []trucks.Domain{}, 0, err
	}

	var totalData int64
	err = repository.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []trucks.Domain{}, 0, err
	}

	var result []trucks.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}
