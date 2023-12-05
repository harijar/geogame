package auth

import (
	"context"
	"crypto/rand"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
)

type Auth struct {
	tokensRepo repo.Tokens
	usersRepo  repo.Users
	logger     *zap.Logger
}

func New(tokensRepo repo.Tokens, usersRepo repo.Users, logger *zap.Logger) *Auth {
	return &Auth{
		tokensRepo: tokensRepo,
		usersRepo:  usersRepo,
		logger:     logger,
	}
}

func (a *Auth) GenerateToken(user *users.User) (string, error) {
	token := make([]byte, 64)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	err = a.tokensRepo.Set(context.Background(), user.ID, string(token))
	return string(token), err
}

func (a *Auth) Register(ctx context.Context, user *users.User) error {
	err := a.usersRepo.UpdateOrSave(ctx, user)
	return err
}
