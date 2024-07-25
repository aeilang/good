package server

import (
	"context"
	"database/sql"
	"time"

	"github.com/aeilang/backend/internal/store/userstore"
	"github.com/aeilang/backend/token"
	"github.com/aeilang/backend/utils"
)

type Auth struct {
	db        *sql.DB
	userStore userstore.Querier
	token     token.Tokener
}

func NewAuth(db *sql.DB, userStore userstore.Querier, token token.Tokener) *Auth {
	return &Auth{
		db:        db,
		userStore: userStore,
		token:     token,
	}
}

type Auther interface {
}

var _ Auther = (*Auth)(nil)

func (a *Auth) GetUserByEmail(ctx context.Context, email string) (userstore.User, error) {
	return a.userStore.GetUserByEmail(ctx, email)
}

type TokenPayload struct {
	Email    string
	Role     string
	Duration time.Duration
}

func (a *Auth) GenerateTokens(access, refresh TokenPayload) (accessToken, refreshToken string, err error) {
	accessToken, err = a.token.CreateAccessToken(token.Data{
		Email:    access.Email,
		Role:     access.Email,
		Duration: access.Duration,
	})
	if err != nil {
		return
	}

	refreshToken, err = a.token.CreateRefreshToken(token.Data{
		Email:    refresh.Email,
		Role:     refresh.Role,
		Duration: refresh.Duration,
	})
	return
}

func (a *Auth) UpdateRefreshToken(ctx context.Context, email, refreshToken string) error {
	return a.userStore.UpdateRefreshTokenByEmail(ctx, userstore.UpdateRefreshTokenByEmailParams{
		RefreshToken: sql.NullString{
			String: refreshToken,
			Valid:  true,
		},
		Email: email,
	})
}

type UserPayload struct {
	Name     string
	Email    string
	Role     string
	Password string
}

func (a *Auth) InsertUser(ctx context.Context, u UserPayload) error {
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return a.userStore.CreateUser(ctx, userstore.CreateUserParams{
		Email:    u.Email,
		Role:     u.Role,
		Name:     u.Name,
		Password: hash,
	})
}

func (a *Auth) ChangePassword(ctx context.Context, password string, email string) error {
	hash, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	err = a.userStore.UpdatePasswordByEmail(ctx, userstore.UpdatePasswordByEmailParams{
		Password: hash,
		Email:    email,
	})

	return err
}
