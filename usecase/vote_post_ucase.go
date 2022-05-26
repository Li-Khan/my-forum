package usecase

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
)

type votePostUsecase struct {
	votePostRepo domain.VotePostRepository
}

func NewVotePostUsecase(v domain.VotePostRepository) domain.VotePostUsecase {
	return &votePostUsecase{
		votePostRepo: v,
	}
}

func (v *votePostUsecase) Like(ctx context.Context, vp *domain.VotePost) (err error) {
	return nil
}

func (v *votePostUsecase) Dislike(ctx context.Context, vp *domain.VotePost) (err error) {
	return nil
}
