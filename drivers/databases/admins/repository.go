package admins

import (
	"context"
	"go-drop-logistik/modules/admins"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreAdminRepository struct {
	tx *gorm.DB
}

func NewPostgreAdminRepository(tx *gorm.DB) admins.Repository {
	return &postgreAdminRepository{
		tx: tx,
	}
}

func (repository *postgreAdminRepository) GetByID(ctx context.Context, id int) (admins.Domain, error) {
	adminsById := Admins{}
	result := repository.tx.Where("admins.id = ?", id).First(&adminsById)
	if result.Error != nil {
		log.Println("[error] admins.repository.GetByID : failed to execute get data admin query", result.Error)
		return admins.Domain{}, result.Error
	}

	return *adminsById.ToDomain(), nil
}

func (repository *postgreAdminRepository) GetByEmail(ctx context.Context, email string) (admins.Domain, error) {
	rec := Admins{}

	err := repository.tx.Where("admins.email = ?", email).First(&rec).Error
	if err != nil {
		log.Println("[error] admins.repository.GetByEmail : failed to execute get data admin query", err)
		return admins.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repository *postgreAdminRepository) Register(ctx context.Context, adminDomain *admins.Domain) error {
	rec := fromDomain(*adminDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] admins.repository.Register : failed to execute register admin query", result.Error)
		return result.Error
	}
	return nil
}
