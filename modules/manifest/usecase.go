package manifest

import (
	"context"
	"errors"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/constants"
	"go-drop-logistik/helpers"
	"time"
)

type ManifestUsecase struct {
	manifestRepository Repository
	contextTimeout     time.Duration
	jwtAuth            *middleware.ConfigJWT
}

func NewManifestUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &ManifestUsecase{
		manifestRepository: ur,
		jwtAuth:            jwtauth,
		contextTimeout:     timeout,
	}
}

func (usecase *ManifestUsecase) StoreManifest(ctx context.Context, manifestDomain *Domain) error {
	manifestDomain.Code = helpers.GenerateManifest()
	manifestDomain.Status = constants.PROCESS

	err := usecase.manifestRepository.StoreManifest(ctx, manifestDomain)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ManifestUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	manifest, err := usecase.manifestRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return manifest, nil
}

func (usecase *ManifestUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := usecase.manifestRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (usecase *ManifestUsecase) Delete(ctx context.Context, id int) error {
	err := usecase.manifestRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ManifestUsecase) Update(ctx context.Context, manifestDomain *Domain, id int) error {
	switch manifestDomain.Status {
		case constants.PROCESS:
			manifestDomain.Status = constants.PROCESS
		case constants.SUCCESS:
			manifestDomain.Status = constants.SUCCESS
		case constants.SHIPPING:
			manifestDomain.Status = constants.SHIPPING
		default:
			return errors.New("message not found")
	}
	
	err := usecase.manifestRepository.Update(ctx, manifestDomain, id)
	if err != nil {
		return err
	}

	return nil
}
