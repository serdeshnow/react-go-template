package repository

import (
	"context"
	"react-go-template/internal/models"
)

type UserRepo interface {
	Get(ctx context.Context, id int) (*models.GetUser, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, userCreate models.CreateUser) (int, error)
	GetPwdByEmail(ctx context.Context, email string) (string, error)
	GetIDByEmail(ctx context.Context, email string) (int, error)
}
