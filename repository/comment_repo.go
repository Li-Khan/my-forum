package repository

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type commentRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) domain.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) Create(ctx context.Context, comment *domain.Comment) (id int64, err error) {
	return 0, nil
}

func (c *commentRepository) Update(ctx context.Context, comment *domain.Comment) (err error) {
	return nil
}

func (c *commentRepository) GetByPostID(ctx context.Context, id int64) (comments []*domain.CommentRequestDTO, err error) {
	return nil, nil
}

func (c *commentRepository) GetByUserID(ctx context.Context, id int64) (comments []*domain.CommentRequestDTO, err error) {
	return nil, nil
}

func (c *commentRepository) Delete(ctx context.Context, id int64) (err error) {
	return nil
}
