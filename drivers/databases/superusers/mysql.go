package superusers

import (
	"context"
	"go-drop-logistik/business/superusers"

	"gorm.io/gorm"
)

type mysqlSuperusersRepository struct {
	Conn *gorm.DB
}

func NewMySQLSuperusersRepository(conn *gorm.DB) superusers.Repository {
	return &mysqlSuperusersRepository{
		Conn: conn,
	}
}

func (repository *mysqlSuperusersRepository) GetByID(ctx context.Context, id int) (superusers.Domain, error) {
	superusersById := Superusers{}
	result := repository.Conn.Where("superusers.id = ?", id).First(&superusersById)
	if result.Error != nil {
		return superusers.Domain{}, result.Error
	}

	return *superusersById.ToDomain(), nil
}

func (repository *mysqlSuperusersRepository) GetByEmail(ctx context.Context, email string) (superusers.Domain, error) {
	rec := Superusers{}

	err := repository.Conn.Where("superusers.email = ?", email).First(&rec).Error
	if err != nil {
		return superusers.Domain{}, err
	}

	return *rec.ToDomain(), nil
}

func (repository *mysqlSuperusersRepository) Register(ctx context.Context, userDomain *superusers.Domain) error {
	rec := fromDomain(*userDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
