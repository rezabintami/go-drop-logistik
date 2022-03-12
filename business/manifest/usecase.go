package manifest

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helper/code"
	"go-drop-logistik/helper/enum"
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
	manifestDomain.Code = code.GenerateManifest()
	manifestDomain.Status = enum.PROCESS

	err := usecase.manifestRepository.StoreManifest(ctx, manifestDomain)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *ManifestUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	users, err := usecase.manifestRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return users, nil
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
	err := usecase.manifestRepository.Update(ctx, manifestDomain, id)
	if err != nil {
		return err
	}

	return nil
}