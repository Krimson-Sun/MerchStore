package mappers

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"merch-store/internal/domain"
	desc "merch-store/pkg/merch"
)

func ItemToProto(item domain.Item) *desc.Item {
	return &desc.Item{
		Id:          item.ID.String(),
		Name:        item.Name,
		Description: item.Description,
		ImageUrl:    item.ImageURL,
		Price:       int32(item.Price),
		InStock:     int32(item.InStock),
		CreatedAt:   timestamppb.New(item.CreatedAt),
		UpdatedAt:   timestamppb.New(item.UpdatedAt),
	}
}
