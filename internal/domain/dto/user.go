package dto

import (
	"merch-store/internal/utils"
)

type CreateUserDTO struct {
	Email    string
	Password string

	FirstName utils.Nullable[string]
	LastName  utils.Nullable[string]
}

type UpdateUserDTO struct {
	FirstName utils.Nullable[string]
	LastName  utils.Nullable[string]
}
