package user

import (
	"context"
	"fmt"

	"merch-store/internal/app/interceptors"
	"merch-store/internal/app/mappers"
	"merch-store/internal/domain"
	"merch-store/internal/domain/dto"
	"merch-store/internal/logger"
	"merch-store/internal/utils"

	desc "merch-store/pkg/merch"

	"github.com/opentracing/opentracing-go"
)

func (i *Implementation) UpdateUser(ctx context.Context, in *desc.UpdateUserRequest) (*desc.UserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.user.Update")
	defer span.Finish()

	if err := in.Validate(); err != nil {
		return nil, fmt.Errorf("%w: %w", domain.ErrInvalidArgument, err)
	}

	id, ok := interceptors.GetUserID(ctx)
	if !ok {
		logger.Errorf("user id not found in context")
		return nil, domain.ErrUnauthorized
	}

	var input dto.UpdateUserDTO
	{
		input.FirstName = utils.NewNullable(in.GetFirstName(), in.GetFirstName() != "")
		input.LastName = utils.NewNullable(in.GetLastName(), in.GetLastName() != "")
	}

	user, err := i.service.UpdateUser(ctx, id, input)
	if err != nil {
		return nil, err
	}

	return &desc.UserResponse{
		User: mappers.UserToProto(user),
	}, nil
}
