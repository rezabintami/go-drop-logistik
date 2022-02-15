package agents

import (
	"context"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business"
	"go-drop-logistik/helper/encrypt"
	"go-drop-logistik/helper/logging"
	"strings"
	"time"
)

type AgentUsecase struct {
	agentRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
	logger         logging.Logger
}

func NewAgentUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &AgentUsecase{
		agentRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
		logger:         logger,
	}
}

func (uc *AgentUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	request := map[string]interface{}{
		"email": email,
		"sso":   sso,
	}

	existedUser, err := uc.agentRepository.GetByEmail(ctx, email)
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

func (uc *AgentUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	request := map[string]interface{}{
		"id": id,
	}

	users, err := uc.agentRepository.GetByID(ctx, id)

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

func (uc *AgentUsecase) Register(ctx context.Context, userDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	request := map[string]interface{}{
		"email": userDomain.Email,
		"name":  userDomain.Name,
	}

	existedUser, err := uc.agentRepository.GetByEmail(ctx, userDomain.Email)
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
		userDomain.Password, _ = encrypt.Hash(userDomain.Password)
	}

	err = uc.agentRepository.Register(ctx, userDomain)
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
