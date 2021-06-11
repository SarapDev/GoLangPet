package account

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository 	Repository
	logger 		log.Logger
}

func NewService(rep Repository, log log.Logger) Service {
	return &service{
		repository: rep,
		logger: log,
	}
}

func (s service) CreateUser(context context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	userUuid, _ := uuid.NewV4()
	id := userUuid.String()
	user := User {
		ID: id,
		Email: email,
		Password: password,
	}

	if err := s.repository.CreateUser(context, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create user", id)
	return "Success", nil
}

func (s service) GetUser(context context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(context, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get User", id)

	return email, nil
}