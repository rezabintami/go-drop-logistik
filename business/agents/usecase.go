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
	jwtusecaseth        *middleware.ConfigJWT
	logger         logging.Logger
}

func NewAgentUsecase(ur Repository, jwtusecaseth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &AgentUsecase{
		agentRepository: ur,
		jwtusecaseth:        jwtusecaseth,
		contextTimeout: timeout,
		logger:         logger,
	}
}

func (usecase *AgentUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	request := map[string]interface{}{
		"email": email,
		"sso":   sso,
	}

	existedUser, err := usecase.agentRepository.GetByEmail(ctx, email)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		usecase.logger.LogEntry(request, result).Error(err.Error())
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", business.ErrEmailPasswordNotFound
	}

	token := usecase.jwtusecaseth.GenerateToken(existedUser.ID, existedUser.Roles)
	result := map[string]interface{}{
		"success": "true",
	}
	usecase.logger.LogEntry(request, result).Info("incoming request")
	return token, nil
}

func (usecase *AgentUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	request := map[string]interface{}{
		"id": id,
	}

	users, err := usecase.agentRepository.GetByID(ctx, id)

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

func (usecase *AgentUsecase) Register(ctx context.Context, agentDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	request := map[string]interface{}{
		"email": agentDomain.Email,
		"name":  agentDomain.Name,
	}

	agentDomain.Roles = "AGENT"
	
	existedUser, err := usecase.agentRepository.GetByEmail(ctx, agentDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"success": "false",
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
		agentDomain.Password, _ = encrypt.Hash(agentDomain.Password)
	}

	err = usecase.agentRepository.Register(ctx, agentDomain)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		usecase.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	usecase.logger.LogEntry(request, result).Info("incoming request")

	return nil
}

func (usecase *AgentUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := usecase.agentRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (usecase *AgentUsecase) Update(ctx context.Context, userDomain *Domain, id int) error {
	err := usecase.agentRepository.Update(ctx, userDomain, id)
	if err != nil {
		return err
	}
	return nil
}