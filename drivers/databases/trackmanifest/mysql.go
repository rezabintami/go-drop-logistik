package trackmanifest

import (
	"context"
	"go-drop-logistik/business/trackmanifest"

	"gorm.io/gorm"
)

type mysqlTrackManifestRepository struct {
	Conn *gorm.DB
}

func NewMySQLTrackManifestRepository(conn *gorm.DB) trackmanifest.Repository {
	return &mysqlTrackManifestRepository{
		Conn: conn,
	}
}

func (repository *mysqlTrackManifestRepository) Store(ctx context.Context, manifestId, trackId int) error {
	trackManifest := &trackmanifest.Domain{
		ManifestID: manifestId,
		TrackID:    trackId,
	}

	rec := fromDomain(*trackManifest)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *mysqlTrackManifestRepository) GetByManifestID(ctx context.Context, id int) (trackmanifest.Domain, error) {
	trackManifest := TrackManifest{}
	result := repository.Conn.Where("manifest_id = ?", id).First(&trackManifest)
	if result.Error != nil {
		return trackmanifest.Domain{}, result.Error
	}

	return *trackManifest.ToDomain(), nil
}

func (repository *mysqlTrackManifestRepository) DeleteByManifest(ctx context.Context, manifestId int) error {
	trackManifest := TrackManifest{}
	result := repository.Conn.Preload("Track").Preload("Manifest").Where("manifest_id = ?", manifestId).Delete(&trackManifest)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
