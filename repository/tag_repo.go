package repository

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type tagRepository struct {
	db *pgxpool.Pool
}

func NewTagRepository(db *pgxpool.Pool) domain.TagRepository {
	return &tagRepository{
		db: db,
	}
}

func (t *tagRepository) Create(ctx context.Context, tag *domain.Tag) (id int64, err error) {
	return 0, nil
}
