package auth

import (
	"context"
	"crypto/rand"
	"github.com/google/uuid"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
)

type Auth struct {
	tokensRepo repo.Tokens
	usersRepo  repo.Users
	logger     *zap.Logger
}

const tokenLenght = 64

func New(tokensRepo repo.Tokens, usersRepo repo.Users, logger *zap.Logger) *Auth {
	return &Auth{
		tokensRepo: tokensRepo,
		usersRepo:  usersRepo,
		logger:     logger,
	}
}

func (a *Auth) GenerateToken() (string, error) {
	token := make([]byte, tokenLenght)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return string(token), err
}

func (a *Auth) RegisterOrUpdate(ctx context.Context, user *users.User) error {
	return a.usersRepo.UpdateOrSave(ctx, user)
}

func (a *Auth) GetUserID(ctx context.Context, token string) (int, error) {
	return a.tokensRepo.GetUserID(ctx, token)
}

func (a *Auth) GetGameID(ctx context.Context, token string) (uuid.UUID, error) {
	return a.tokensRepo.GetGameID(ctx, token)
}

func (a *Auth) SetUserID(ctx context.Context, token string, id int) error {
	return a.tokensRepo.SetUserID(ctx, token, id)
}

func (a *Auth) SetGameID(ctx context.Context, token string, id uuid.UUID) error {
	return a.tokensRepo.SetGameID(ctx, token, id)
}

func (a *Auth) UpdateUser(ctx context.Context, user *users.User) error {
	return a.usersRepo.Update(ctx, user)
}
