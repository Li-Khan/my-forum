package domain

import (
	"context"
	"time"
)

// User ...
type User struct {
	UserID    int64     `json:"user_id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (id int64, err error)
	Update(ctx context.Context, user *User) (err error)
	GetByID(ctx context.Context, id int64) (user *User, err error)
	Delete(ctx context.Context, id int64) (err error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (id int64, err error)
	Update(ctx context.Context, user *User) (err error)
	GetByID(ctx context.Context, id int64) (user *User, err error)
	Delete(ctx context.Context, id int64) (err error)
}
