package admins

import (
	"context"
	"go-drop-logistik/business/admins"

	"github.com/jinzhu/gorm"
)

type mysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMySQLAdminRepository(conn *gorm.DB) admins.Repository {
	return &mysqlAdminRepository{
		Conn: conn,
	}
}

func (repository *mysqlAdminRepository) GetByID(ctx context.Context, id int) (admins.Domain, error) {
	adminsById := Admins{}
	result := repository.Conn.Where("admins.id = ?", id).First(&adminsById)
	if result.Error != nil {
		return admins.Domain{}, result.Error
	}

	return *adminsById.ToDomain(), nil
}

func (repository *mysqlAdminRepository) GetByEmail(ctx context.Context, email string) (admins.Domain, error) {
	rec := Admins{}

	err := repository.Conn.Where("admins.email = ?", email).First(&rec).Error
	if err != nil {
		return admins.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repository *mysqlAdminRepository) Register(ctx context.Context, adminDomain *admins.Domain) error {
	rec := fromDomain(*adminDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
