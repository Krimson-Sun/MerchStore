package service

import (
	"context"
	"time"

	"merch-store/internal/domain"
)

type jwtProvider interface {
	GeneratePair(ctx context.Context, userID, pairID domain.ID, atTime time.Time) (domain.Tokens, error)
	VerifyPair(ctx context.Context, userID domain.ID, tokens domain.Tokens, atTime time.Time) error
	ParseToken(ctx context.Context, token string) (domain.ID, error)
}

type sessionRepository interface {
	GetSessionByToken(ctx context.Context, token string) (domain.Session, error)
	SetSessionExpired(ctx context.Context, id domain.ID, expiredAt time.Time) error
	CreateSession(ctx context.Context, session domain.Session) (domain.Session, error)
}

type userRepository interface {
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	GetUserByID(ctx context.Context, id domain.ID) (domain.User, error)
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) (domain.User, error)
}

type unitOfWork interface {
	Begin(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	InTransaction(ctx context.Context, f func(ctx context.Context) error) error
}

type Service struct {
	jwtProvider       jwtProvider
	sessionRepository sessionRepository
	userRepository    userRepository
	unitOfWork        unitOfWork
}

func New(
	unitOfWork unitOfWork,
	jwtProvider jwtProvider,
	sessionRepository sessionRepository,
	userRepository userRepository,
) *Service {
	return &Service{
		unitOfWork:        unitOfWork,
		jwtProvider:       jwtProvider,
		sessionRepository: sessionRepository,
		userRepository:    userRepository,
	}
}
