package tracks

import (
	"context"
	"go-drop-logistik/modules/tracks"
	"log"
	"time"

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

func (repository *postgreTrackRepository) Delete(ctx context.Context, trackId int) error {
	result := repository.tx.Preload("Agent").Where("tracks.id = ?", trackId).Delete(Tracks{})
	if result.Error != nil {
		log.Println("[error] tracks.repository.Delete : failed to execute delete track query", result.Error)
		return result.Error
	}
	return nil
}

func (repository *postgreTrackRepository) Update(ctx context.Context, trackId int, trackDomain *tracks.Domain) error {
	result := repository.tx.Exec(
		"UPDATE tracks SET message = ?, destination_agent_id = ?, current_agent_id = ?, updated_at = ? WHERE id = ?",
		trackDomain.Message, trackDomain.DestinationAgentID, trackDomain.CurrentAgentID, time.Now(), trackId)
	if result.Error != nil {
		log.Println("[error] tracks.repository.Update : failed to execute update manifest query", result.Error)
		return result.Error
	}
	return nil
}
