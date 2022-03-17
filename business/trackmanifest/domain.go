package trackmanifest

import (
	"context"
	"go-drop-logistik/business/manifest"
	"go-drop-logistik/business/tracks"
)

type Domain struct {
	TrackID    int
	Track      *tracks.Domain
	ManifestID int
	Manifest   *manifest.Domain
}

type Usecase interface {
	Store(ctx context.Context, manifestId, trackId int) error
	GetByManifestID(ctx context.Context, id int) (Domain, error)
	DeleteByManifest(ctx context.Context, manifestId int) error
}

type Repository interface {
	Store(ctx context.Context, manifestId, trackId int) error
	GetByManifestID(ctx context.Context, id int) (Domain, error)
	DeleteByManifest(ctx context.Context, manifestId int) error
}
