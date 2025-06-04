package items

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"merch-store/internal/app/mappers"
	"merch-store/internal/domain"
	"merch-store/internal/logger"
	desc "merch-store/pkg/merch"
)

func (i *Implementation) GetItem(ctx context.Context, in *desc.GetItemRequest) (*desc.Item, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.items.GetItem")
	defer span.Finish()

	span.SetTag("item_id", in.GetItemId())
	if err := in.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err)
		return nil, fmt.Errorf("%w: %w", domain.ErrInvalidArgument, err)
	}

	id, err := domain.ParseID(in.GetItemId())
	if err != nil {
		logger.Errorf("error parsing item id: %v", err)
		return nil, fmt.Errorf("%w: %w", domain.ErrInvalidArgument, err)
	}

	item, err := i.service.GetItemByID(ctx, id)
	if err != nil {
		logger.Errorf("error getting item: %v", err)
		return nil, err
	}

	return mappers.ItemToProto(item), nil
}
