package tracks

import (
	"context"
	"errors"
	"fmt"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helper/enum"
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
	if trackDomain.Message == fmt.Sprintf("%s %s", enum.TRACKING_PROCESS_MESSAGE, agentName) {
		trackDomain.Message = fmt.Sprintf("%s %s", enum.TRACKING_PROCESS_MESSAGE, agentName)
	} else if trackDomain.Message == enum.TRACKING_SHIPPING_MESSAGE {
		trackDomain.Message = enum.TRACKING_SHIPPING_MESSAGE
	} else if trackDomain.Message == enum.TRACKING_SUCCESS_MESSAGE {
		trackDomain.Message = enum.TRACKING_SUCCESS_MESSAGE
	} else {
		return 0, errors.New("message not found")
	}

	id, err := usecase.trackRepository.StoreTrack(ctx, trackDomain)
	if err != nil {
		return 0, err
	}

	return id, nil
}
