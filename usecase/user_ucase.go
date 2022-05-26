package usecase

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: u,
	}
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) (id int64, err error) {
	return 0, nil
}

func (u *userUsecase) Update(ctx context.Context, user *domain.User) (err error) {
	return nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (user *domain.User, err error) {
	return nil, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int64) (err error) {
	return nil
}