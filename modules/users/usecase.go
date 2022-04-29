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

func (uc *UserUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	existedUser, err := uc.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !helpers.ValidateHash(password, existedUser.Password) && !sso {
		return "", helpers.ErrEmailPasswordNotFound
	}

	token := uc.jwtAuth.GenerateToken(existedUser.ID, existedUser.Name, existedUser.Roles)

	return token, nil
}

func (uc *UserUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	users, err := uc.userRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return users, nil
}

func (uc *UserUsecase) Register(ctx context.Context, userDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByEmail(ctx, userDomain.Email)
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

	err = uc.userRepository.Register(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}
