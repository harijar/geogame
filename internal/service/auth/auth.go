package auth

import (
	"context"
	"crypto/rand"
	"github.com/google/uuid"
	"github.com/harijar/geogame/internal/repo"
	"github.com/harijar/geogame/internal/repo/postgres/users"
	"go.uber.org/zap"
	"time"
)

const (
	tokenLenght = 64
	userIDTTL   = 720 * time.Hour
	gameIDTTL   = 24 * time.Hour
)

type Auth struct {
	redisRepo repo.Redis
	usersRepo repo.Users
	logger    *zap.Logger
}

func New(redisRepo repo.Redis, usersRepo repo.Users, logger *zap.Logger) *Auth {
	return &Auth{
		redisRepo: redisRepo,
		usersRepo: usersRepo,
		logger:    logger,
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

func (a *Auth) UserExists(ctx context.Context, id int) (bool, error) {
	return a.usersRepo.Exists(ctx, id)
}

func (a *Auth) RegisterOrUpdate(ctx context.Context, user *users.User) error {
	return a.usersRepo.UpdateOrSave(ctx, user)
}

func (a *Auth) GetUserID(ctx context.Context, token string) (int, error) {
	return a.redisRepo.GetUserID(ctx, token)
}

func (a *Auth) GetGameID(ctx context.Context, token string) (uuid.UUID, error) {
	return a.redisRepo.GetGameID(ctx, token)
}

func (a *Auth) SetUserID(ctx context.Context, token string, id int) error {
	return a.redisRepo.SetUserID(ctx, token, id, userIDTTL)
}

func (a *Auth) SetGameID(ctx context.Context, token string, id uuid.UUID) error {
	return a.redisRepo.SetGameID(ctx, token, id, gameIDTTL)
}
