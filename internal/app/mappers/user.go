package mappers

import (
	"merch-store/internal/domain"
	desc "merch-store/pkg/merch"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToProto(user domain.User) *desc.User {
	userProto := &desc.User{
		Id:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}

	return userProto
}
