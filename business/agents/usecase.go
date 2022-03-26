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
	contextTimeout  time.Duration
	jwtusecaseth    *middleware.ConfigJWT
	logger          logging.Logger
}

func NewAgentUsecase(ur Repository, jwtusecaseth *middleware.ConfigJWT, timeout time.Duration, logger logging.Logger) Usecase {
	return &AgentUsecase{
		agentRepository: ur,
		jwtusecaseth:    jwtusecaseth,
		contextTimeout:  timeout,
		logger:          logger,
	}
}

func (usecase *AgentUsecase) Login(ctx context.Context, email, password string, sso bool) (string, error) {
	existedUser, err := usecase.agentRepository.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, existedUser.Password) && !sso {
		return "", business.ErrEmailPasswordNotFound
	}

	token := usecase.jwtusecaseth.GenerateToken(existedUser.ID, existedUser.Name, existedUser.Roles)

	return token, nil
}

func (usecase *AgentUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	users, err := usecase.agentRepository.GetByID(ctx, id)

	if err != nil {
		return Domain{}, err
	}

	return users, nil
}

func (usecase *AgentUsecase) Register(ctx context.Context, agentDomain *Domain, sso bool) error {
	ctx, cancel := context.WithTimeout(ctx, usecase.contextTimeout)
	defer cancel()

	agentDomain.Roles = "AGENT"

	existedUser, err := usecase.agentRepository.GetByEmail(ctx, agentDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (ExistingDomain{}) {
		return business.ErrDuplicateData
	}

	if !sso {
		agentDomain.Password, _ = encrypt.Hash(agentDomain.Password)
	}

	err = usecase.agentRepository.Register(ctx, agentDomain)
	if err != nil {
		return err
	}

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
