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

func (au *AgentUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	request := map[string]interface{}{
		"email": email,
		"sso":   sso,
	}

	existedUser, err := au.agentRepository.GetByEmail(ctx, email)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		au.logger.LogEntry(request, result).Error(err.Error())
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", business.ErrEmailPasswordNotFound
	}

	token := au.jwtAuth.GenerateToken(existedUser.ID, existedUser.Roles)
	result := map[string]interface{}{
		"success": "true",
	}
	au.logger.LogEntry(request, result).Info("incoming request")
	return token, nil
}

func (au *AgentUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	request := map[string]interface{}{
		"id": id,
	}

	users, err := au.agentRepository.GetByID(ctx, id)

	if err != nil {
		result := map[string]interface{}{
			"error": err.Error(),
		}
		au.logger.LogEntry(request, result).Error(err.Error())
		return Domain{}, err
	}

	result := map[string]interface{}{
		"id":    users.ID,
		"name":  users.Name,
		"email": users.Email,
	}

	au.logger.LogEntry(request, result).Info("incoming request")

	return users, nil
}

func (au *AgentUsecase) Register(ctx context.Context, agentDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()

	request := map[string]interface{}{
		"email": agentDomain.Email,
		"name":  agentDomain.Name,
	}

	agentDomain.Roles = "AGENT"
	
	existedUser, err := au.agentRepository.GetByEmail(ctx, agentDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			result := map[string]interface{}{
				"success": "false",
				"error":   err.Error(),
			}
			au.logger.LogEntry(request, result).Error(err.Error())
			return err
		}
	}
	if existedUser != (Domain{}) {
		return business.ErrDuplicateData
	}

	if !sso {
		agentDomain.Password, _ = encrypt.Hash(agentDomain.Password)
	}

	err = au.agentRepository.Register(ctx, agentDomain)
	if err != nil {
		result := map[string]interface{}{
			"success": "false",
			"error":   err.Error(),
		}
		au.logger.LogEntry(request, result).Error(err.Error())
		return err
	}

	result := map[string]interface{}{
		"success": "true",
	}
	au.logger.LogEntry(request, result).Info("incoming request")

	return nil
}

func (au *AgentUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 25
	}

	res, total, err := au.agentRepository.Fetch(ctx, page, perpage)
	if err != nil {
		return []Domain{}, 0, err
	}

	return res, total, nil
}

func (au *AgentUsecase) Update(ctx context.Context, userDomain *Domain, id int) error {
	err := au.agentRepository.Update(ctx, userDomain, id)
	if err != nil {
		return err
	}
	return nil
}