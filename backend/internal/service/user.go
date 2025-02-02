package service

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"react-fsd-template/internal/models"
	"react-fsd-template/internal/repository"
	"react-fsd-template/pkg/log"
)

type userService struct {
	userRepo repository.UserRepo
	logger   *log.Logs
}

func InitUserService(userRepo repository.UserRepo, logger *log.Logs) User {
	return userService{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u userService) GetMe(ctx context.Context, id int) (*models.GetUser, error) {
	user, err := u.userRepo.Get(ctx, id)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}

	return user, nil
}

func (u userService) Delete(ctx context.Context, id int) error {
	err := u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (u userService) Create(ctx context.Context, user models.CreateUser) (int, error) {
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 11)
	user.Password = string(hashedPwd)

	id, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u userService) Login(ctx context.Context, user models.CreateUser) (int, error) {

	hashedPwd, err := u.userRepo.GetPwdByEmail(ctx, user.Email)
	if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(user.Password))
	if err != nil {
		return 0, err
	}

	var id int
	id, err = u.userRepo.GetIDByEmail(ctx, user.Email)
	if err != nil {
		return 0, err
	}

	return id, nil
}
