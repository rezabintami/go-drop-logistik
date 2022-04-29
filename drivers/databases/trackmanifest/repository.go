package trackmanifest

import (
	"context"
	"go-drop-logistik/modules/trackmanifest"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreTrackManifestRepository struct {
	tx *gorm.DB
}

func NewPostgreTrackManifestRepository(tx *gorm.DB) trackmanifest.Repository {
	return &postgreTrackManifestRepository{
		tx: tx,
	}
}

func (repository *postgreTrackManifestRepository) Store(ctx context.Context, manifestId, trackId int) error {
	trackManifest := &trackmanifest.Domain{
		ManifestID: manifestId,
		TrackID:    trackId,
	}

	rec := fromDomain(*trackManifest)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] trackmanifests.repository.Store : failed to execute store trackmanifest query", result.Error)
		return result.Error
	}

	return nil
}

func (repository *postgreTrackManifestRepository) GetByManifestID(ctx context.Context, id int) (trackmanifest.Domain, error) {
	trackManifest := TrackManifest{}
	result := repository.tx.Where("manifest_id = ?", id).First(&trackManifest)
	if result.Error != nil {
		log.Println("[error] trackmanifests.repository.GetByManifestID : failed to execute get data trackmanifest query", result.Error)
		return trackmanifest.Domain{}, result.Error
	}

	return *trackManifest.ToDomain(), nil
}

func (repository *postgreTrackManifestRepository) GetAllByManifestID(ctx context.Context, id int) ([]trackmanifest.Domain, error) {
	allTrackManifest := []TrackManifest{}

	result := repository.tx.Preload("Track.StartAgent").Preload("Track.CurrentAgent").Preload("Track.DestinationAgent").Where("manifest_id = ?", id).Find(&allTrackManifest)
	if result.Error != nil {
		log.Println("[error] trackmanifests.repository.GetAllByManifestID : failed to execute get data trackmanifests query", result.Error)
		return []trackmanifest.Domain{}, result.Error
	}

	allTrackManifestDomain := []trackmanifest.Domain{}
	for _, value := range allTrackManifest {
		allTrackManifestDomain = append(allTrackManifestDomain, *value.ToDomain())
	}

	return allTrackManifestDomain, nil
}
func (repository *postgreTrackManifestRepository) DeleteByManifest(ctx context.Context, manifestId int) error {
	trackManifest := TrackManifest{}
	result := repository.tx.Preload("Track").Preload("Manifest").Where("manifest_id = ?", manifestId).Delete(&trackManifest)
	if result.Error != nil {
		log.Println("[error] trackmanifests.repository.DeleteByManifest : failed to execute delete trackmanifest query", result.Error)
		return result.Error
	}

	return nil
}
