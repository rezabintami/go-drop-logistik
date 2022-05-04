package trackmanifest

import (
	"go-drop-logistik/drivers/databases/manifest"
	"go-drop-logistik/drivers/databases/tracks"
	"go-drop-logistik/modules/trackmanifest"
)

type TrackManifest struct {
	ID         int `gorm:"primary_key"`
	TrackID    int
	Track      *tracks.Tracks `gorm:"foreignkey:TrackID;references:ID"`
	ManifestID int
	Manifest   *manifest.Manifest `gorm:"foreignkey:ManifestID;references:ID"`
}

func fromDomain(trackManifestDomain trackmanifest.Domain) *TrackManifest {
	return &TrackManifest{
		ID:         trackManifestDomain.ID,
		TrackID:    trackManifestDomain.TrackID,
		ManifestID: trackManifestDomain.ManifestID,
	}
}

func (rec *TrackManifest) ToDomain() *trackmanifest.Domain {
	return &trackmanifest.Domain{
		TrackID:    rec.TrackID,
		Track:      rec.Track.ToDomain(),
		ManifestID: rec.ManifestID,
		Manifest:   rec.Manifest.ToDomain(),
	}
}
