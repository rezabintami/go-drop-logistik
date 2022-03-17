package tracks

import (
	"context"
	"go-drop-logistik/app/middleware"
	"time"
)

type TrackUsecase struct {
	trackRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewTrackUsecase(repo Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &TrackUsecase{
		trackRepository: repo,
		jwtAuth:         jwtauth,
		contextTimeout:  timeout,
	}
}

func (usecase *TrackUsecase) StoreTrack(ctx context.Context, trackDomain *Domain) error {
	err := usecase.trackRepository.StoreTrack(ctx, trackDomain)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrackUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	users, err := usecase.trackRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return users, nil
}