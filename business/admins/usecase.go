package admins

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business"
	"go-drop-logistik/helper/encrypt"
	"go-drop-logistik/helper/logging"
	"strings"
	"time"
)

type AdminUsecase struct {
	adminRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
	logger          logging.Logger
}

func NewAdminUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &AdminUsecase{
		adminRepository: ur,
		jwtAuth:         jwtauth,
		contextTimeout:  timeout,
		logger:          logger,
	}
}

func (usecase *AdminUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	request := map[string]interface{}{
		"email": email,
		"sso":   sso,
	}

	existedUser, err := usecase.adminRepository.GetByEmail(ctx, email)
	if err != nil {
		result := map[string]interface{}{
			"susecasecess": "false",
			"error":   err.Error(),
		}
		usecase.logger.LogEntry(request, result).Error(err.Error())
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", business.ErrEmailPasswordNotFound
	}

	token := usecase.jwtAuth.GenerateToken(existedUser.ID, existedUser.Roles)
	result := map[string]interface{}{
		"susecasecess": "true",
	}
	usecase.logger.LogEntry(request, result).Info("incoming request")
	return token, nil
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
				"error":   err.Error(),
			}
			usecase.logger.LogEntry(request, result).Error(err.Error())
			return err
		}
	}
	if existedUser != (Domain{}) {
		return business.ErrDuplicateData
	}

	if !sso {
		adminDomain.Password, _ = encrypt.Hash(adminDomain.Password)
	}

	err = usecase.adminRepository.Register(ctx, adminDomain)
	if err != nil {
		result := map[string]interface{}{
			"susecasecess": "false",
			"error":   err.Error(),
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
