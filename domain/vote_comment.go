package domain

import "context"

type VoteComment struct {
	VotePostID int64 `json:"vote_post_id,omitempty"`
	UserID     int64 `json:"user_id,omitempty"`
	PostID     int64 `json:"post_id,omitempty"`
	Vote       int   `json:"vote,omitempty"`
}

type VoteCommentUsecase interface {
	Like(ctx context.Context, vp *VoteComment) (err error)
	Dislike(ctx context.Context, vp *VoteComment) (err error)
}

type VoteCommentRepository interface {
	Like(ctx context.Context, vp *VoteComment) (err error)
	Dislike(ctx context.Context, vp *VoteComment) (err error)
}
