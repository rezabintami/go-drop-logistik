package manifest

import (
	"context"
	"go-drop-logistik/modules/manifest"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type postgreManifestRepository struct {
	tx *gorm.DB
}

func NewPostgreManifestRepository(tx *gorm.DB) manifest.Repository {
	return &postgreManifestRepository{
		tx: tx,
	}
}

func (repository *postgreManifestRepository) StoreManifest(ctx context.Context, manifestDomain *manifest.Domain) error {
	rec := fromDomain(*manifestDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] manifests.repository.StoreManifest : failed to execute store manifest query", result.Error)
		return result.Error
	}
	return nil
}

func (repository *postgreManifestRepository) GetByID(ctx context.Context, id int) (manifest.Domain, error) {
	rec := Manifest{}
	result := repository.tx.Preload("Driver").Preload("Driver.Truck").Where("id = ?", id).First(&rec)
	if result.Error != nil {
		log.Println("[error] manifests.repository.GetByID : failed to execute get data manifest query", result.Error)
		return manifest.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}

func (repository *postgreManifestRepository) Fetch(ctx context.Context, page, perpage int) ([]manifest.Domain, int, error) {
	rec := []Manifest{}

	offset := (page - 1) * perpage
	err := repository.tx.Preload("Driver").Preload("Driver.Truck").Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		log.Println("[error] manifests.repository.Fetch : failed to execute fetch manifests query", err)
		return []manifest.Domain{}, 0, err
	}

	var totalData int64
	err = repository.tx.Model(&rec).Count(&totalData).Error
	if err != nil {
		log.Println("[error] manifests.repository.Fetch : failed to execute count manifests query", err)

		return []manifest.Domain{}, 0, err
	}

	var result []manifest.Domain
	for _, value := range rec {
		result = append(result, *value.ToDomain())
	}

	return result, int(totalData), nil
}

func (repository *postgreManifestRepository) Delete(ctx context.Context, id int) error {
	rec := Manifest{}
	result := repository.tx.Where("id = ?", id).Delete(&rec)
	if result.Error != nil {
		log.Println("[error] manifests.repository.Delete : failed to execute delete manifest query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreManifestRepository) Update(ctx context.Context, manifestDomain *manifest.Domain, id int) error {
	result := repository.tx.Exec(
		"UPDATE manifests SET status = ?, driver_id = ?, updated_at = ? WHERE id = ?",
		manifestDomain.Status, manifestDomain.DriverID, time.Now(), id)
	if result.Error != nil {
		log.Println("[error] manifests.repository.Update : failed to execute update manifest query", result.Error)
		return result.Error
	}

	return nil
}
