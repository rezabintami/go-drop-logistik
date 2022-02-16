package superusers

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business"
	"go-drop-logistik/helper/encrypt"
	"go-drop-logistik/helper/logging"
	"strings"
	"time"
)

type SuperuserUsecase struct {
	superuserRepository Repository
	contextTimeout      time.Duration
	jwtAuth             *middleware.ConfigJWT
	logger              logging.Logger
}

func NewSuperuserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &SuperuserUsecase{
		superuserRepository: ur,
		jwtAuth:             jwtauth,
		contextTimeout:      timeout,
		logger:              logger,
	}
}

func (uc *SuperuserUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	request := map[string]interface{}{
		"email": email,
		"sso":   sso,
	}

	existedUser, err := uc.superuserRepository.GetByEmail(ctx, email)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		uc.logger.LogEntry(request, result).Error(err.Error())
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", business.ErrEmailPasswordNotFound
	}

	token := uc.jwtAuth.GenerateToken(existedUser.ID, existedUser.Roles)
	result := map[string]interface{}{
		"success": "true",
	}
	uc.logger.LogEntry(request, result).Info("incoming request")
	return token, nil
}

func (uc *SuperuserUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	request := map[string]interface{}{
		"id": id,
	}

	users, err := uc.superuserRepository.GetByID(ctx, id)

	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		uc.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"id":    users.ID,
		"name":  users.Name,
		"email": users.Email,
	}

	uc.logger.LogEntry(request, result).Info("incoming request")

	return users, nil
}

func (uc *SuperuserUsecase) Register(ctx context.Context, superuserDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	request := map[string]interface{}{
		"email": superuserDomain.Email,
		"name":  superuserDomain.Name,
	}

	superuserDomain.Roles = "SUPERUSER"

	existedUser, err := uc.superuserRepository.GetByEmail(ctx, superuserDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"success": "false",
				"error":   err.Error(),
			}
			uc.logger.LogEntry(request, result).Error(err.Error())
			return err
		}
	}
	if existedUser != (Domain{}) {
		return business.ErrDuplicateData
	}

	if !sso {
		superuserDomain.Password, _ = encrypt.Hash(superuserDomain.Password)
	}

	err = uc.superuserRepository.Register(ctx, superuserDomain)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		uc.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	uc.logger.LogEntry(request, result).Info("incoming request")

	return nil
}
