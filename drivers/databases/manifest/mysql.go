package manifest

import (
	"context"
	"go-drop-logistik/business/manifest"

	"gorm.io/gorm"
)

type mysqlManifestRepository struct {
	Conn *gorm.DB
}

func NewMySQLManifestRepository(conn *gorm.DB) manifest.Repository {
	return &mysqlManifestRepository{
		Conn: conn,
	}
}


func (repository *mysqlManifestRepository) StoreManifest(ctx context.Context, receiptDomain *manifest.Domain) error {
	rec := fromDomain(*receiptDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *mysqlManifestRepository) GetByID(ctx context.Context, id int) (manifest.Domain, error) {
	rec := Manifest{}
	result := repository.Conn.Where("id = ?", id).First(&rec)
	if result.Error != nil {
		return manifest.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}


func (repository *mysqlManifestRepository) Fetch(ctx context.Context, page, perpage int) ([]manifest.Domain, int, error) {
	rec := []Manifest{}

	offset := (page - 1) * perpage
	err := repository.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []manifest.Domain{}, 0, err
	}

	var totalData int64
	err = repository.Conn.Model(&rec).Count(&totalData).Error
	if err != nil {
		return []manifest.Domain{}, 0, err
	}

	var result []manifest.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}

func (repository *mysqlManifestRepository) Delete(ctx context.Context, id int) error {
	rec := Manifest{}
	result := repository.Conn.Where("id = ?", id).Delete(&rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
