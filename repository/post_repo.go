package repository

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) domain.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) Create(ctx context.Context, post *domain.Post) (id int64, err error) {
	return 0, nil
}

func (p *postRepository) Update(ctx context.Context, post *domain.Post) (err error) {
	return nil
}

func (p *postRepository) GetAll(ctx context.Context) (posts *[]domain.PostRequestDTO, err error) {
	return nil, nil
}

func (p *postRepository) GetByID(ctx context.Context, id int64) (post *domain.PostRequestDTO, err error) {
	return nil, nil
}

func (p *postRepository) Delete(ctx context.Context, id int64) (err error) {
	return nil
}
