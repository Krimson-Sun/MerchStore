package user

import (
	"context"
	"fmt"

	"merch-store/internal/app/mappers"
	"merch-store/internal/domain"
	"merch-store/internal/domain/dto"
	"merch-store/internal/logger"
	"merch-store/internal/utils"

	desc "merch-store/pkg/merch"

	"github.com/opentracing/opentracing-go"
)

func (i *Implementation) CreateUser(ctx context.Context, in *desc.CreateUserRequest) (*desc.UserResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "api.user.Create")
	defer span.Finish()

	if err := in.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err)
		return nil, fmt.Errorf("%w: %w", domain.ErrInvalidArgument, err)
	}

	var input dto.CreateUserDTO
	{
		input.Email = in.Email
		input.Password = in.Password

		input.FirstName = utils.NewNullable(in.GetFirstName(), in.GetFirstName() != "")
		input.LastName = utils.NewNullable(in.GetLastName(), in.GetLastName() != "")
	}

	user, err := i.service.CreateUser(ctx, input)
	if err != nil {
		logger.Errorf("error creating user: %v", err)
		return nil, err
	}

	return &desc.UserResponse{
		User: mappers.UserToProto(user),
	}, nil
}
