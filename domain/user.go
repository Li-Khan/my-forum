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

type UserRequestDTO struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserUsecase interface {
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, id int64) error
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Delete(ctx context.Context, id int64) error
}
