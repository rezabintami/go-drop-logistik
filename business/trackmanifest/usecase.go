package trackmanifest

import (
	"context"
	"go-drop-logistik/app/middleware"
	"time"
)

type TrackManifestUsecase struct {
	trackManifestRepository Repository
	contextTimeout          time.Duration
	jwtAuth                 *middleware.ConfigJWT
}

func NewTrackManifestUsecase(repo Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &TrackManifestUsecase{
		trackManifestRepository: repo,
		jwtAuth:                 jwtauth,
		contextTimeout:          timeout,
	}
}

func (usecase *TrackManifestUsecase) Store(ctx context.Context, manifestId, trackId int) error {
	err := usecase.trackManifestRepository.Store(ctx, manifestId, trackId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrackManifestUsecase) GetByManifestID(ctx context.Context, id int) (Domain, error) {
	res, err := usecase.trackManifestRepository.GetByManifestID(ctx, id)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (usecase *TrackManifestUsecase) GetAllByManifestID(ctx context.Context, id int) ([]Domain, error) {
	res, err := usecase.trackManifestRepository.GetAllByManifestID(ctx, id)

	if err != nil {
		return []Domain{}, err
	}

	return res, nil
}

func (usecase *TrackManifestUsecase) DeleteByManifest(ctx context.Context, manifestId int) error {
	err := usecase.trackManifestRepository.DeleteByManifest(ctx, manifestId)
	if err != nil {
		return err
	}

	return nil
}
