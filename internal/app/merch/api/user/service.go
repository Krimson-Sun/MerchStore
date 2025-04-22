package user

import (
	"context"
	"merch-store/internal/domain"
	"merch-store/internal/domain/dto"
	desc "merch-store/pkg/merch"
)

type Service interface {
	CreateUser(ctx context.Context, dto dto.CreateUserDTO) (domain.User, error)
	GetUserByID(ctx context.Context, id domain.ID) (domain.User, error)
	UpdateUser(ctx context.Context, id domain.ID, dto dto.UpdateUserDTO) (domain.User, error)
}

type Implementation struct {
	service Service
	desc.UnimplementedUserServiceServer
}

func New(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
