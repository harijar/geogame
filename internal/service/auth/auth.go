package auth

import (
	"context"
	"crypto/rand"
	"github.com/harijar/geogame/internal/repo"
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

func (a *Auth) GetTokenAndSave(id int, firstName, lastName, username string) (string, error) {
	token := make([]byte, 64)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	err = a.tokensRepo.Set(context.Background(), id, string(token))
	if err != nil {
		return "", err
	}
	err = a.usersRepo.Save(context.Background(), id, firstName, lastName, username)
	return string(token), err
}
