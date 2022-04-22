package phones

import (
	"context"
	"go-drop-logistik/business/phones"

	"github.com/jinzhu/gorm"
)

type mysqlPhoneRepository struct {
	Conn *gorm.DB
}

func NewMySQLPhoneRepository(conn *gorm.DB) phones.Repository {
	return &mysqlPhoneRepository{
		Conn: conn,
	}
}

func (repository *mysqlPhoneRepository) StorePhone(ctx context.Context, phoneDomain *phones.Domain) (int, error) {
	rec := fromDomain(*phoneDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return 0, result.Error
	}

	return rec.ID, nil
}

func (repository *mysqlPhoneRepository) GetByID(ctx context.Context, id int) (phones.Domain, error) {
	phone := Phones{}
	result := repository.Conn.Where("id = ?", id).First(&phone)
	if result.Error != nil {
		return phones.Domain{}, result.Error
	}

	return *phone.ToDomain(), nil
}

func (repository *mysqlPhoneRepository) GetAll(ctx context.Context) ([]phones.Domain, error) {
	rec := []Phones{}
	result := repository.Conn.Find(&rec)
	if result.Error != nil {
		return []phones.Domain{}, result.Error
	}
	phoneDomain := []phones.Domain{}
	for _, value := range rec {
		phoneDomain = append(phoneDomain, *value.ToDomain())
	}

	return phoneDomain, nil
}


func (repository *mysqlPhoneRepository) Update(ctx context.Context, phoneDomain *phones.Domain, id int) error {
	phoneUpdate := fromDomain(*phoneDomain)

	result := repository.Conn.Where("id = ?", id).Updates(&phoneUpdate)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlPhoneRepository) Delete(ctx context.Context, id int) error {
	phoneDelete := Phones{}
	result := repository.Conn.Where("id = ?", id).Delete(&phoneDelete)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
