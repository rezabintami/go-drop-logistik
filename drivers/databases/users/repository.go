package users

import (
	"context"
	"go-drop-logistik/business/users"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreUsersRepository struct {
	tx *gorm.DB
}

func NewPostgreUsersRepository(tx *gorm.DB) users.Repository {
	return &postgreUsersRepository{
		tx: tx,
	}
}

func (repository *postgreUsersRepository) GetByID(ctx context.Context, id int) (users.Domain, error) {
	usersById := Users{}
	result := repository.tx.Where("users.id = ?", id).First(&usersById)
	if result.Error != nil {
		log.Println("[error] users.repository.GetByID : failed to execute get data user query", result.Error)
		return users.Domain{}, result.Error
	}

	return *usersById.ToDomain(), nil
}

func (repository *postgreUsersRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := Users{}

	err := repository.tx.Where("users.email = ?", email).First(&rec).Error
	if err != nil {
		log.Println("[error] users.repository.GetByEmail : failed to execute get data user query", err)
		return users.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repository *postgreUsersRepository) Register(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] users.repository.Register : failed to execute register user query", result.Error)
		return result.Error
	}
	return nil
}
