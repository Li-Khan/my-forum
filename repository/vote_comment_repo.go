package repository

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type voteCommentRepository struct {
	db *pgxpool.Pool
}

func NewVoteCommentRepository(db *pgxpool.Pool) domain.VoteCommentRepository {
	return &voteCommentRepository{
		db: db,
	}
}

func (v *voteCommentRepository) Like(ctx context.Context, vp *domain.VoteComment) (err error) {
	return nil
}

func (v *voteCommentRepository) Dislike(ctx context.Context, vp *domain.VoteComment) (err error) {
	return nil
}
