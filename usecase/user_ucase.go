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
	id, err = u.userRepo.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) Update(ctx context.Context, user *domain.User) (err error) {
	return nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (user *domain.User, err error) {
	return nil, nil
}

func (u *userUsecase) GetByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	user, err = u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) Delete(ctx context.Context, id int64) (err error) {
	return nil
}
