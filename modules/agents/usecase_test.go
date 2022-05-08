package agents_test

import (
	"context"
	"errors"
	"go-drop-logistik/app/middleware"
	"go-drop-logistik/helpers"
	"go-drop-logistik/modules/agents"
	agentMock "go-drop-logistik/modules/agents/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	agentRepository agentMock.Repository
	agentUsecase    agents.Usecase
	jwtAuth         *middleware.ConfigJWT
)

func setup() {
	jwtAuth = &middleware.ConfigJWT{SecretJWT: "abc123", ExpiresDuration: 2}
	agentUsecase = agents.NewAgentUsecase(&agentRepository, jwtAuth, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestLoginAgent(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		email := "agent@gmail.com"
		sso := false
		pass, _ := helpers.Hash("123123")
		agentDomain := agents.ExistingDomain{
			ID:       1,
			Password: pass,
			Name:     "agent",
			Email:    email,
			Roles:    "AGENT",
		}

		agentRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(agentDomain, nil).Once()

		accessToken, refreshToken, err := agentUsecase.Login(context.Background(), email, "123123", sso)
		assert.Nil(t, err)
		assert.NotEmpty(t, accessToken)
		assert.NotEmpty(t, refreshToken)
	})
	t.Run("test case 2, password error", func(t *testing.T) {
		pass, _ := helpers.Hash("1231231")
		agentDomain := agents.ExistingDomain{
			ID:       1,
			Password: pass,
			Name:     "agent",
			Email:    "agent@gmail.com",
			Roles:    "AGENT",
		}

		agentRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(agentDomain, nil).Once()

		_, _, err := agentUsecase.Login(context.Background(), "agent@gmail.com", "123123", false)
		assert.Equal(t, err, helpers.ErrEmailPasswordNotFound)

	})

	t.Run("test case 3, error record", func(t *testing.T) {

		errRepository := errors.New("error record")
		agentRepository.On("GetByEmail", mock.Anything, mock.AnythingOfType("string")).Return(agents.ExistingDomain{}, errRepository).Once()

		accessToken, refreshToken, err := agentUsecase.Login(context.Background(), "logistik@gmail.com", "123123", false)

		assert.Equal(t, err, errRepository)
		assert.Equal(t, "", accessToken)
		assert.Equal(t, "", refreshToken)
	})
}

func TestRegisterAgent(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := agents.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "agent",
			Email:    "agent@gmail.com",
		}
		agentRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(agents.ExistingDomain{}, nil).Once()
		agentRepository.On("Register", mock.Anything, mock.Anything).Return(nil).Once()

		err := agentUsecase.Register(context.Background(), &domain, false)

		assert.Nil(t, err)
	})

	t.Run("test case 2, duplicate data", func(t *testing.T) {
		domain := agents.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "agent",
			Email:    "agent@gmail.com",
		}

		domainDuplicate := agents.ExistingDomain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "agent",
			Email:    "agent@gmail.com",
		}

		errRepository := errors.New("duplicate data")
		agentRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(domainDuplicate, errRepository).Once()

		err := agentUsecase.Register(context.Background(), &domain, false)

		assert.Equal(t, err, helpers.ErrDuplicateData)
	})

	t.Run("test case 3, data has exist", func(t *testing.T) {
		domain := agents.Domain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "agent",
			Email:    "agent@gmail.com",
		}

		domainDuplicate := agents.ExistingDomain{
			ID:       1,
			Password: "asyudasd820aisd",
			Name:     "agent",
			Email:    "agent@gmail.com",
		}

		agentRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(domainDuplicate, nil).Once()

		err := agentUsecase.Register(context.Background(), &domain, false)

		assert.Equal(t, err, helpers.ErrDuplicateData)
	})

	t.Run("test case 4, register failed", func(t *testing.T) {
		domain := agents.Domain{
			ID:    1,
			Name:  "agent",
			Email: "agent@gmail.com",
		}
		errRepository := errors.New("register failed")
		agentRepository.On("GetByEmail", mock.Anything, mock.Anything).Return(agents.ExistingDomain{}, nil).Once()
		agentRepository.On("Register", mock.Anything, mock.Anything).Return(errRepository).Once()

		err := agentUsecase.Register(context.Background(), &domain, false)

		assert.Equal(t, err, errRepository)
	})
}

func TestUpdateAgent(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := agents.Domain{
			ID:        1,
			Password:  "asyudasd820aisd",
			Name:      "agent",
			Email:     "agent@gmail.com",
			Address:   "Jl. Kebon Jeruk",
			Balance:   20000,
			Latitude:  5.8234324,
			Longitude: -5.8234324,
		}
		agentRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := agentUsecase.Update(context.Background(), &domain, 1)

		assert.Nil(t, err)
	})

	t.Run("test case 2, id not found", func(t *testing.T) {
		domain := agents.Domain{
			ID:        1,
			Password:  "asyudasd820aisd",
			Name:      "agent",
			Email:     "agent@gmail.com",
			Address:   "Jl. Kebon Jeruk",
			Balance:   20000,
			Latitude:  5.8234324,
			Longitude: -5.8234324,
		}
		errRepository := errors.New("id not found")
		agentRepository.On("Update", mock.Anything, mock.Anything, mock.AnythingOfType("int")).Return(errRepository).Once()

		err := agentUsecase.Update(context.Background(), &domain, -1)

		assert.Equal(t, err, errRepository)
	})
}

func TestFetchDataAgent(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := []agents.Domain{
			{
				ID:        1,
				Name:      "agent",
				Email:     "agent@gmail.com",
				Address:   "Jl. Kebon Jeruk",
				Latitude:  5.8234324,
				Longitude: -5.8234324,
			},
			{
				ID:        2,
				Name:      "agent2",
				Email:     "agent2@gmail.com",
				Address:   "Jl. Kebon Selatan",
				Latitude:  5.8234324,
				Longitude: -5.8234324,
			},
		}

		agentRepository.On("Fetch", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, 2, nil).Once()

		result, total, err := agentUsecase.Fetch(context.Background(), 1, 2)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, total, 2)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("data not found")

		agentRepository.On("Fetch", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return([]agents.Domain{}, 0, errRepository).Once()

		result, _, err := agentUsecase.Fetch(context.Background(), 1, 2)

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})

	t.Run("test case 3, fetching without params", func(t *testing.T) {
		domain := []agents.Domain{
			{
				ID:        1,
				Name:      "agent",
				Email:     "agent@gmail.com",
				Address:   "Jl. Kebon Jeruk",
				Latitude:  5.8234324,
				Longitude: -5.8234324,
			},
			{
				ID:        2,
				Name:      "agent2",
				Email:     "agent2@gmail.com",
				Address:   "Jl. Kebon Selatan",
				Latitude:  5.8234324,
				Longitude: -5.8234324,
			},
		}

		agentRepository.On("Fetch", mock.Anything, mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(domain, 2, nil).Once()

		result, total, err := agentUsecase.Fetch(context.Background(), 0, 0)

		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
		assert.Equal(t, total, 2)
	})
}

func TestGetByIdAgent(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		domain := agents.Domain{
			ID:    1,
			Name:  "agent",
			Email: "agent@gmail.com",
		}
		agentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(domain, nil).Once()

		result, err := agentUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, domain.ID, result.ID)
	})

	t.Run("test case 2, data not found", func(t *testing.T) {
		errRepository := errors.New("data not found")
		agentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(agents.Domain{}, errRepository).Once()
		result, err := agentUsecase.GetByID(context.Background(), -1)
		assert.Equal(t, result, agents.Domain{})
		assert.Equal(t, err, errRepository)
	})
}

func TestDeleteAgent(t *testing.T) {
	t.Run("test case 1, valid test", func(t *testing.T) {
		agentRepository.On("Delete", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := agentUsecase.Delete(context.Background(), 1)

		assert.Nil(t, err)
	})
}
