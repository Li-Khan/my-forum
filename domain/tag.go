package domain

import "context"

type Tag struct {
	TagID int64  `json:"tag_id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type TagUsecase interface {
	Create(ctx context.Context, tag *Tag) (id int64, err error)
}

type TagRepository interface {
	Create(ctx context.Context, tag *Tag) (id int64, err error)
}
