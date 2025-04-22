package domain

import (
	"time"

	"github.com/google/uuid"
)

type ID uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func (i ID) String() string {
	return uuid.UUID(i).String()
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ID{}, err
	}
	return ID(id), nil
}

type Model struct {
	ID        ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewModel() Model {
	return Model{
		ID:        NewID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type User struct {
	Model

	Email     string
	Password  string
	FirstName string
	LastName  string
}

func NewUser(
	Email string,
	Password string,
	FirstName string,
	LastName string,
) User {
	return User{
		Model:     NewModel(),
		Email:     Email,
		Password:  Password,
		FirstName: FirstName,
		LastName:  LastName,
	}
}

type Session struct {
	Model

	UserID    ID
	ExpiredAt time.Time
	Token     string
}

func NewSession(userID ID, expiredAt time.Time, token string) Session {
	return Session{
		Model:     NewModel(),
		UserID:    userID,
		ExpiredAt: expiredAt,
		Token:     token,
	}
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}
