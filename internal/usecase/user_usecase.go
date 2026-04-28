package usecase

import (
	"context"
	"go-starter/internal/model"
	"go-starter/internal/repository"
)

// UserUsecase adalah interface untuk logika bisnis terkait User
type UserUsecase interface {
	GetProfile(ctx context.Context, id int) (*model.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

// NewUserUsecase untuk inisialisasi
func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) GetProfile(ctx context.Context, id int) (*model.User, error) {
	user, err := u.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}
