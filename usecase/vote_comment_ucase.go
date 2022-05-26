package usecase

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
)

type voteCommentUsecase struct {
	voteCommentRepo domain.VoteCommentRepository
}

func NewVoteCommentUsecase(v domain.VoteCommentRepository) domain.VoteCommentUsecase {
	return &voteCommentUsecase{
		voteCommentRepo: v,
	}
}

func (v *voteCommentUsecase) Like(ctx context.Context, vp *domain.VoteComment) (err error)
func (v *voteCommentUsecase) Dislike(ctx context.Context, vp *domain.VoteComment) (err error)
