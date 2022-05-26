package domain

import (
	"context"
	"time"
)

// Post ...
type Post struct {
	PostID    int64     `json:"post_id,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// PostRequestDTO ...
type PostRequestDTO struct {
	PostID    int64     `json:"post_id,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Title     string    `json:"title,omitempty"`
	Text      string    `json:"text,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// PostResponseDTO ...
type PostResponseDTO struct {
	PostID    int64     `json:"post_id,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Text      string    `json:"text,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type PostUsecase interface {
	Create(ctx context.Context, post *Post) (id int64, err error)
	Update(ctx context.Context, post *Post) (err error)
	GetAll(ctx context.Context) (posts *[]PostRequestDTO, err error)
	GetByID(ctx context.Context, id int64) (post *PostRequestDTO, err error)
	Delete(ctx context.Context, id int64) (err error)
}

type PostRepository interface {
	Create(ctx context.Context, post *Post) (id int64, err error)
	Update(ctx context.Context, post *Post) (err error)
	GetAll(ctx context.Context) (posts *[]PostRequestDTO, err error)
	GetByID(ctx context.Context, id int64) (post *PostRequestDTO, err error)
	Delete(ctx context.Context, id int64) (err error)
}
