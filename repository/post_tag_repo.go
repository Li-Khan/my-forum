package repository

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postTagRepository struct {
	db *pgxpool.Pool
}

func NewPostTagRepository(db *pgxpool.Pool) domain.PostTagRepository {
	return &postTagRepository{
		db: db,
	}
}

func (p *postTagRepository) Create(ctx context.Context, pt *domain.PostTag) (id int64, err error) {
	return 0, nil
}
