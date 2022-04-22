package tracks

import (
	"context"
	"go-drop-logistik/business/tracks"

	"github.com/jinzhu/gorm"
)

type mysqlTrackRepository struct {
	Conn *gorm.DB
}

func NewMySQLTrackRepository(conn *gorm.DB) tracks.Repository {
	return &mysqlTrackRepository{
		Conn: conn,
	}
}

func (repository *mysqlTrackRepository) StoreTrack(ctx context.Context, trackDomain *tracks.Domain) (int, error) {
	rec := fromDomain(*trackDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return 0, result.Error
	}
	return rec.ID, nil
}
