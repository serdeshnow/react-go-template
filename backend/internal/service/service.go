package service

import (
	"context"
	"react-fsd-template/internal/models"
)

type User interface {
	GetMe(ctx context.Context, id int) (*models.GetUser, error)
	Delete(ctx context.Context, id int) error
	Create(ctx context.Context, user models.CreateUser) (int, error)
	Login(ctx context.Context, user models.CreateUser) (int, error)
}
