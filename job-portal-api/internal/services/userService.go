package services

import (
	"context"
	"job-portal-api/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

//go:generate mockgen -source service.go -destination mockmodels/service_mock.go -package mockmodels

type Service interface {
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
		error)
	CreateCompany(ctx context.Context, nu models.NewCompany, UserId int) (models.Company, error)
	ViewCompanyAll(ctx context.Context, UserId string) ([]models.Company, error)
	ViewJob(ctx context.Context, UserId string) ([]models.Job, error)
	AutoMigrate() error
}

type Store struct {
	Service
}

func NewStore(s Service) Store {
	return Store{Service: s}
}
