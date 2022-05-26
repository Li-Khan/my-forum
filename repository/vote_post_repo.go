package repository

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type votePostRepository struct {
	db *pgxpool.Pool
}

func NewVotePostRepository(db *pgxpool.Pool) domain.VotePostRepository {
	return &votePostRepository{
		db: db,
	}
}

func (v *votePostRepository) Like(ctx context.Context, vp *domain.VotePost) (err error) {
	return nil
}

func (v *votePostRepository) Dislike(ctx context.Context, vp *domain.VotePost) (err error) {
	return nil
}
