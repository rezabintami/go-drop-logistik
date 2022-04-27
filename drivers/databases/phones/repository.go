package phones

import (
	"context"
	"go-drop-logistik/business/phones"
	"log"

	"github.com/jinzhu/gorm"
)

type postgrePhoneRepository struct {
	tx *gorm.DB
}

func NewpostgrePhoneRepository(tx *gorm.DB) phones.Repository {
	return &postgrePhoneRepository{
		tx: tx,
	}
}

func (repository *postgrePhoneRepository) StorePhone(ctx context.Context, phoneDomain *phones.Domain) (int, error) {
	rec := fromDomain(*phoneDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] phones.repository.StorePhone : failed to execute store phone query", result.Error)
		return 0, result.Error
	}

	return rec.ID, nil
}

func (repository *postgrePhoneRepository) GetByID(ctx context.Context, id int) (phones.Domain, error) {
	phone := Phones{}
	result := repository.tx.Where("id = ?", id).First(&phone)
	if result.Error != nil {
		log.Println("[error] phones.repository.GetByID : failed to execute get data phone query", result.Error)
		return phones.Domain{}, result.Error
	}

	return *phone.ToDomain(), nil
}

func (repository *postgrePhoneRepository) GetAll(ctx context.Context) ([]phones.Domain, error) {
	rec := []Phones{}
	result := repository.tx.Find(&rec)
	if result.Error != nil {
		log.Println("[error] phones.repository.GetAll : failed to execute get data phones query", result.Error)
		return []phones.Domain{}, result.Error
	}
	phoneDomain := []phones.Domain{}
	for _, value := range rec {
		phoneDomain = append(phoneDomain, *value.ToDomain())
	}

	return phoneDomain, nil
}

func (repository *postgrePhoneRepository) Update(ctx context.Context, phoneDomain *phones.Domain, id int) error {
	phoneUpdate := fromDomain(*phoneDomain)

	result := repository.tx.Where("id = ?", id).Updates(&phoneUpdate)
	if result.Error != nil {
		log.Println("[error] phones.repository.Update : failed to execute update phone query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgrePhoneRepository) Delete(ctx context.Context, id int) error {
	phoneDelete := Phones{}
	result := repository.tx.Where("id = ?", id).Delete(&phoneDelete)
	if result.Error != nil {
		log.Println("[error] phones.repository.Delete : failed to execute delete phone query", result.Error)
		return result.Error
	}

	return nil
}
