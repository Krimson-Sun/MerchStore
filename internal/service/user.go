package service

import (
	"context"
	"time"

	"merch-store/internal/domain"
	"merch-store/internal/domain/dto"
	"merch-store/internal/utils"

	"github.com/opentracing/opentracing-go"
)

func (s *Service) CreateUser(ctx context.Context, dto dto.CreateUserDTO) (domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.CreateUser")
	defer span.Finish()

	hashedPass, err := utils.HashPassword(dto.Password)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.NewUser(
		dto.Email,
		hashedPass,
		dto.FirstName.V,
		dto.LastName.V,
	)

	err = s.unitOfWork.InTransaction(ctx, func(ctx context.Context) error {
		user, err = s.userRepository.CreateUser(ctx, user)
		return err
	})
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id domain.ID) (domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.GetUserByID")
	defer span.Finish()

	return s.userRepository.GetUserByID(ctx, id)
}

func (s *Service) UpdateUser(ctx context.Context, id domain.ID, dto dto.UpdateUserDTO) (domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.UpdateUser")
	defer span.Finish()

	user, err := s.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	{
		if dto.LastName.IsValid {
			user.LastName = dto.LastName.V
		}

		if dto.FirstName.IsValid {
			user.FirstName = dto.FirstName.V
		}

		user.UpdatedAt = time.Now()
	}

	err = s.unitOfWork.InTransaction(ctx, func(ctx context.Context) error {
		user, err = s.userRepository.UpdateUser(ctx, user)

		return err
	})
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
