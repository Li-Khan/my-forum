package domain

import (
	"context"
	"time"
)

// Comment ...
type Comment struct {
	CommentID int64     `json:"comment_id,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	PostID    int64     `json:"post_id,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// CommentRequestDTO ...
type CommentRequestDTO struct {
	CommentID int64     `json:"comment_id,omitempty"`
	UserID    int64     `json:"user_id,omitempty"`
	PostID    int64     `json:"post_id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CommentUsecase interface {
	Create(ctx context.Context, c *Comment) (id int64, err error)
	Update(ctx context.Context, c *Comment) (err error)
	GetByPostID(ctx context.Context, id int64) (comments []*CommentRequestDTO, err error)
	GetByUserID(ctx context.Context, id int64) (comments []*CommentRequestDTO, err error)
	Delete(ctx context.Context, id int64) (err error)
}

type CommentRepository interface {
	Create(ctx context.Context, c *Comment) (id int64, err error)
	Update(ctx context.Context, c *Comment) (err error)
	GetByPostID(ctx context.Context, id int64) (comments []*CommentRequestDTO, err error)
	GetByUserID(ctx context.Context, id int64) (comments []*CommentRequestDTO, err error)
	Delete(ctx context.Context, id int64) (err error)
}
