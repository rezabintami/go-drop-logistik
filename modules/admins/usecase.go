package admins

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helpers"
	"strings"
	"time"
)

type AdminUsecase struct {
	adminRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
	logger          helpers.Logger
}

func NewAdminUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger helpers.Logger) Usecase {
	return &AdminUsecase{
		adminRepository: ur,
		jwtAuth:         jwtauth,
		contextTimeout:  timeout,
		logger:          logger,
	}
}

func (usecase *AdminUsecase) Login(ctx context.Context, email, password string, sso bool) (string, string, error) {
	request := map[string]interface{}{
		"email": email,
		"sso":   sso,
	}

	existedUser, err := usecase.adminRepository.GetByEmail(ctx, email)
	if err != nil {
		result := map[string]interface{}{
			"susecasecess": "false",
			"error":        err.Error(),
		}
		usecase.logger.LogEntry(request, result).Error(err.Error())
		return "", "", err
	}

	if !helpers.ValidateHash(password, existedUser.Password) && !sso {
		return "", "", helpers.ErrEmailPasswordNotFound
	}

	accessToken := usecase.jwtAuth.GenerateToken(existedUser.ID, existedUser.Name, existedUser.Roles)
	refreshToken := usecase.jwtAuth.GenerateRefreshToken(existedUser.ID)

	return accessToken, refreshToken, nil
}

func (usecase *AdminUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	request := map[string]interface{}{
		"id": id,
	}

	users, err := usecase.adminRepository.GetByID(ctx, id)

	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		usecase.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"id":    users.ID,
		"name":  users.Name,
		"email": users.Email,
	}

	usecase.logger.LogEntry(request, result).Info("incoming request")

	return users, nil
}

func (usecase *AdminUsecase) Register(ctx context.Context, adminDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	request := map[string]interface{}{
		"email": adminDomain.Email,
		"name":  adminDomain.Name,
	}

	adminDomain.Roles = "ADMIN"

	existedUser, err := usecase.adminRepository.GetByEmail(ctx, adminDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"susecasecess": "false",
				"error":        err.Error(),
			}
			usecase.logger.LogEntry(request, result).Error(err.Error())
			return err
		}
	}
	if existedUser != (Domain{}) {
		return helpers.ErrDuplicateData
	}

	if !sso {
		adminDomain.Password, _ = helpers.Hash(adminDomain.Password)
	}

	err = usecase.adminRepository.Register(ctx, adminDomain)
	if err != nil {
		result := map[string]interface{}{
			"susecasecess": "false",
			"error":        err.Error(),
		}
		usecase.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"susecasecess": "true",
	}
	usecase.logger.LogEntry(request, result).Info("incoming request")

	return nil
}
