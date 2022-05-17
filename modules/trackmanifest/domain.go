package trackmanifest

import (
	"context"
	"go-drop-logistik/modules/manifest"
	"go-drop-logistik/modules/tracks"
)

type Domain struct {
	ID         int
	TrackID    int
	Track      *tracks.Domain
	ManifestID int
	Manifest   *manifest.Domain
}

type Usecase interface {
	Store(ctx context.Context, manifestId, trackId int) error
	GetByManifestID(ctx context.Context, manifestId int) (Domain, error)
	GetAllByManifestID(ctx context.Context, manifestId int) ([]Domain, error)
	DeleteByManifest(ctx context.Context, manifestId int) error
	Delete(ctx context.Context, manifestId, trackId int) error
}

type Repository interface {
	Store(ctx context.Context, manifestId, trackId int) error
	GetByManifestID(ctx context.Context, manifestId int) (Domain, error)
	GetAllByManifestID(ctx context.Context, manifestId int) ([]Domain, error)
	DeleteByManifest(ctx context.Context, manifestId int) error
	Delete(ctx context.Context, manifestId, trackId int) error
}
