package repository

import (
	"context"
	"errors"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/Li-Khan/my-forum/lib/mistake"
	"github.com/jackc/pgconn"
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
	stmt := `INSERT INTO "user" (
		"email",
		"username",
		"password",
		"created_at",
		"updated_at"
	) VALUES ($1, $2, $3, $4, $5) RETURNING "user_id"`

	row := u.db.QueryRow(ctx, stmt, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)

	err = row.Scan(&id)
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return 0, mistake.ErrConflict
		}
	}

	if err != nil {
		return 0, err
	}

	return id, nil
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
