package trucks

import (
	"context"
	"go-drop-logistik/modules/trucks"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type postgreTruckRepository struct {
	tx *gorm.DB
}

func NewPostgreTruckRepository(tx *gorm.DB) trucks.Repository {
	return &postgreTruckRepository{
		tx: tx,
	}
}

func (repository *postgreTruckRepository) StoreTruck(ctx context.Context, truckDomain *trucks.Domain) error {
	rec := fromDomain(*truckDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] trucks.repository.StoreTruck : failed to execute store truck query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreTruckRepository) GetByID(ctx context.Context, id int) (trucks.Domain, error) {
	truck := Trucks{}
	result := repository.tx.Where("id = ?", id).First(&truck)
	if result.Error != nil {
		log.Println("[error] trucks.repository.GetByID : failed to execute get data truck query", result.Error)
		return trucks.Domain{}, result.Error
	}

	return *truck.ToDomain(), nil
}

func (repository *postgreTruckRepository) Update(ctx context.Context, truckDomain *trucks.Domain, id int) error {
	truckUpdate := fromDomain(*truckDomain)
	truckUpdate.UpdatedAt = time.Now()

	result := repository.tx.Table("trucks").Where("id = ?", id).Updates(&truckUpdate)
	if result.Error != nil {
		log.Println("[error] trucks.repository.Update : failed to execute update truck query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreTruckRepository) Delete(ctx context.Context, id int) error {
	truckDelete := Trucks{}
	result := repository.tx.Where("id = ?", id).Delete(&truckDelete)
	if result.Error != nil {
		log.Println("[error] trucks.repository.Delete : failed to execute delete truck query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreTruckRepository) Fetch(ctx context.Context, page, perpage int) ([]trucks.Domain, int, error) {
	rec := []Trucks{}

	offset := (page - 1) * perpage
	err := repository.tx.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		log.Println("[error] trucks.repository.Fetch : failed to execute fetch trucks query", err)
		return []trucks.Domain{}, 0, err
	}

	var totalData int64
	err = repository.tx.Model(&rec).Count(&totalData).Error
	if err != nil {
		log.Println("[error] trucks.repository.Fetch : failed to execute count trucks query", err)
		return []trucks.Domain{}, 0, err
	}

	var result []trucks.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}

func (repository *postgreTruckRepository) CheckByID(ctx context.Context, id int) error {
	truck := Trucks{}
	result := repository.tx.Where("id = ?", id).First(&truck)
	if result.Error != nil {
		log.Println("[error] trucks.repository.CheckByID : failed to execute check truck query", result.Error)
		return result.Error
	}

	return nil
}