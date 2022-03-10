package phones

import (
	"context"
	"go-drop-logistik/business/phones"

	"gorm.io/gorm"
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
	result := repository.Conn.Where("phones.id = ?", id).First(&phone)
	if result.Error != nil {
		return phones.Domain{}, result.Error
	}

	return *phone.ToDomain(), nil
}
