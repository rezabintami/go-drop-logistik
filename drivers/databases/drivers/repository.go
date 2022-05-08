package drivers

import (
	"context"
	"errors"
	"go-drop-logistik/modules/drivers"
	"log"
	"time"

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
	driver := Drivers{}
	result := repository.tx.Preload("Truck").Where("id = ?", id).First(&driver)
	if result.Error != nil {
		log.Println("[error] drivers.repository.GetByID : failed to execute get data driver query", result.Error)
		return drivers.Domain{}, result.Error
	}

	return *driver.ToDomain(), nil
}

func (repository *postgreDriverRepository) CheckByID(ctx context.Context, id int) error {
	driverCheck := Drivers{}
	result := repository.tx.Where("id = ?", id).First(&driverCheck)
	if result.Error != nil {
		log.Println("[error] drivers.repository.CheckByID : failed to execute check data driver query", result.Error)
		return result.Error
	}

	return nil
}

// This Gorm same as sql query in below comment
// "UPDATE drivers SET name = 'Dana' , phone = '083123246347' , address = 'Jalan Mangga Manis No 134', truck_id = 2 
// WHERE EXISTS(SELECT * FROM drivers left join trucks on trucks.id = 1 WHERE trucks.id = 1 AND trucks.deleted_at is null) AND drivers.id = 2"
func (repository *postgreDriverRepository) Update(ctx context.Context, driverDomain *drivers.Domain, id int) error {
	driverUpdate := fromDomain(*driverDomain)
	driverUpdate.UpdatedAt = time.Now()
	
	subQuery := repository.tx.Table("drivers").Joins("left join trucks on trucks.id = ?",
		driverDomain.TruckID).Where("trucks.id = ? AND trucks.deleted_at is null AND drivers.id = ?",
		driverDomain.TruckID, id).QueryExpr()
	result := repository.tx.Table("drivers").Where("EXISTS(?) AND drivers.id = ?", subQuery, id).Updates(&driverUpdate)
	if result.RowsAffected == 0 {
		log.Println("[error] drivers.repository.Update : failed to execute update driver query", result.Error)
		return errors.New("truck not found")
	}

	return nil
}

func (repository *postgreDriverRepository) Delete(ctx context.Context, id int) error {
	driverDelete := Drivers{}
	result := repository.tx.Preload("Truck").Where("id = ?", id).Delete(&driverDelete)
	if result.Error != nil {
		log.Println("[error] drivers.repository.Delete : failed to execute delete driver query", result.Error)
		return result.Error
	}

	return nil
}
