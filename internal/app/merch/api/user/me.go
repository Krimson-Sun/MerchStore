package user

import (
	"context"
	"fmt"
	"merch-store/internal/app/interceptors"
	"merch-store/internal/app/mappers"
	"merch-store/internal/domain"
	desc "merch-store/pkg/merch"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) GetMe(ctx context.Context, _ *emptypb.Empty) (*desc.UserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.user.GetMe")
	defer span.Finish()

	id, ok := interceptors.GetUserID(ctx)
	if !ok {
		return nil, fmt.Errorf("user id not found in context: %w", domain.ErrUnauthorized)
	}

	user, err := i.service.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &desc.UserResponse{
		User: mappers.UserToProto(user),
	}, nil
}
