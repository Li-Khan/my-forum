package repository

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) (id int64, err error) {
	return 0, nil
}

func (u *userRepository) Update(ctx context.Context, user *domain.User) (err error) {
	return nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (user *domain.User, err error) {
	return nil, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) (err error) {
	return nil
}
