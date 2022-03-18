package trackmanifest

import (
	"go-drop-logistik/business/trackmanifest"
	"go-drop-logistik/drivers/databases/manifest"
	"go-drop-logistik/drivers/databases/tracks"
)

type TrackManifest struct {
	TrackID    int
	Track      *tracks.Tracks `gorm:"foreignkey:TrackID;references:ID"`
	ManifestID int
	Manifest   *manifest.Manifest `gorm:"foreignkey:ManifestID;references:ID"`
}

func fromDomain(trackManifestDomain trackmanifest.Domain) *TrackManifest {
	return &TrackManifest{
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
