package users

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helpers"
	"strings"
	"time"
)

type UserUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &UserUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (usecase *UserUsecase) Login(ctx context.Context, email, password string, sso bool) (string, string, error) {
	existedUser, err := usecase.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", "", err
	}

	if !helpers.ValidateHash(password, existedUser.Password) && !sso {
		return "", "", helpers.ErrEmailPasswordNotFound
	}

	accessToken := usecase.jwtAuth.GenerateToken(existedUser.ID, existedUser.Name, existedUser.Roles)
	refreshToken := usecase.jwtAuth.GenerateRefreshToken(existedUser.ID)

	return accessToken, refreshToken, nil
}

func (usecase *UserUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	users, err := usecase.userRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return users, nil
}

func (usecase *UserUsecase) Register(ctx context.Context, userDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	existedUser, err := usecase.userRepository.GetByEmail(ctx, userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return helpers.ErrDuplicateData
	}

	if !sso {
		userDomain.Password, _ = helpers.Hash(userDomain.Password)
	}

	err = usecase.userRepository.Register(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}
