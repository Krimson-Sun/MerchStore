package auth

import (
	"context"
	"merch-store/internal/domain"
	desc "merch-store/pkg/merch"
)

type Service interface {
	Login(ctx context.Context, email string, password string) (domain.Tokens, error)
	Logout(ctx context.Context, refreshToken string) error
	Refresh(ctx context.Context, tokens domain.Tokens) (domain.Tokens, error)
	ParseToken(ctx context.Context, token string) (domain.ID, error)
}

type Implementation struct {
	service Service

	desc.UnimplementedAuthServiceServer
}

func New(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
