package tracks

import (
	"context"
	"go-drop-logistik/business/tracks"

	"gorm.io/gorm"
)

type mysqlTrackRepository struct {
	Conn *gorm.DB
}

func NewMySQLTrackRepository(conn *gorm.DB) tracks.Repository {
	return &mysqlTrackRepository{
		Conn: conn,
	}
}

func (repository *mysqlTrackRepository) StoreTrack(ctx context.Context, trackDomain *tracks.Domain) error {
	rec := fromDomain(*trackDomain)

	result := repository.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repository *mysqlTrackRepository) GetByID(ctx context.Context, id int) (tracks.Domain, error) {
	rec := Tracks{}
	result := repository.Conn.Where("id = ?", id).First(&rec)
	if result.Error != nil {
		return tracks.Domain{}, result.Error
	}

	return *rec.ToDomain(), nil
}

