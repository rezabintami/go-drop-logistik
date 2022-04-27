package tracks

import (
	"context"
	"go-drop-logistik/business/tracks"
	"log"

	"github.com/jinzhu/gorm"
)

type postgreTrackRepository struct {
	tx *gorm.DB
}

func NewPostgreTrackRepository(tx *gorm.DB) tracks.Repository {
	return &postgreTrackRepository{
		tx: tx,
	}
}

func (repository *postgreTrackRepository) StoreTrack(ctx context.Context, trackDomain *tracks.Domain) (int, error) {
	rec := fromDomain(*trackDomain)

	result := repository.tx.Create(rec)
	if result.Error != nil {
		log.Println("[error] tracks.repository.StoreTrack : failed to execute store track query", result.Error)
		return 0, result.Error
	}
	return rec.ID, nil
}
