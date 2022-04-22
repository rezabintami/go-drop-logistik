package users

import (
	"context"
	"go-drop-logistik/business/users"

	"github.com/jinzhu/gorm"
)

type mysqlUsersRepository struct {
	Conn *gorm.DB
}

func NewMySQLUserRepository(conn *gorm.DB) users.Repository {
	return &mysqlUsersRepository{
		Conn: conn,
	}
}

func (repository *mysqlUsersRepository) GetByID(ctx context.Context, id int) (users.Domain, error) {
	usersById := Users{}
	result := repository.Conn.Where("users.id = ?", id).First(&usersById)
	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return *usersById.ToDomain(), nil
}

func (repository *mysqlUsersRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := Users{}

	err := repository.Conn.Where("users.email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repository *mysqlUsersRepository) Register(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
