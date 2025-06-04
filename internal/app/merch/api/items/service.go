package items

import (
	"context"
	"merch-store/internal/domain"
	desc "merch-store/pkg/merch"
)

type Service interface {
	GetItemByID(ctx context.Context, id domain.ID) (domain.Item, error)
}

type Implementation struct {
	service Service
	desc.UnimplementedMerchServiceServer
}

func New(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
