package domain

import "context"

type PostTag struct {
	PostTagID int64 `json:"post_tag_id,omitempty"`
	PostID    int64 `json:"post_id,omitempty"`
	TagID     int64 `json:"tag_id,omitempty"`
}

type PostTagUsecase interface {
	Create(ctx context.Context, pt *PostTag) (id int64, err error)
}

type PostTagRepository interface {
	Create(ctx context.Context, pt *PostTag) (id int64, err error)
}
