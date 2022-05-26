package domain

import "context"

type VotePost struct {
	VotePostID int64 `json:"vote_post_id,omitempty"`
	UserID     int64 `json:"user_id,omitempty"`
	PostID     int64 `json:"post_id,omitempty"`
	Vote       int   `json:"vote,omitempty"`
}

type VotePostUsecase interface {
	Like(ctx context.Context, vp *VotePost) (err error)
	Dislike(ctx context.Context, vp *VotePost) (err error)
}

type VotePostRepository interface {
	Like(ctx context.Context, vp *VotePost) (err error)
	Dislike(ctx context.Context, vp *VotePost) (err error)
}
