package tracks

import (
	"context"
	"errors"
	"fmt"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/constants"
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

func (usecase *TrackUsecase) StoreTrack(ctx context.Context, trackDomain *Domain, agentName string) (int, error) {
	fmt.Println("message : ", trackDomain.Message)
	fmt.Printf("%s %s", constants.TRACKING_PROCESS_MESSAGE, agentName)
	switch trackDomain.Message {
		case constants.TRACKING_PROCESS_MESSAGE:
			trackDomain.Message = fmt.Sprintf("%s %s", constants.TRACKING_PROCESS_MESSAGE, agentName)
		case constants.TRACKING_SHIPPING_MESSAGE:
			trackDomain.Message = constants.TRACKING_SHIPPING_MESSAGE
		case constants.TRACKING_SUCCESS_MESSAGE:
			trackDomain.Message = constants.TRACKING_SUCCESS_MESSAGE
		default:
			return 0, errors.New("message not found")
	}

	id, err := usecase.trackRepository.StoreTrack(ctx, trackDomain)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (usecase *TrackUsecase) Delete(ctx context.Context, trackId, agentId int) error {
	err := usecase.trackRepository.Delete(ctx, trackId, agentId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *TrackUsecase) Update(ctx context.Context, trackId, agentId int, trackDomain *Domain) error {
	err := usecase.trackRepository.Update(ctx, trackId, agentId, trackDomain)
	if err != nil {
		return err
	}
	return nil
}
