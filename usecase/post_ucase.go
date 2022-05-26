package usecase

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
)

type postUsecase struct {
	postRepo domain.PostRepository
}

func NewPostUsecase(p domain.PostRepository) domain.PostUsecase {
	return &postUsecase{
		postRepo: p,
	}
}

func (p *postUsecase) Create(ctx context.Context, post *domain.Post) (id int64, err error) {
	return 0, nil
}

func (p *postUsecase) Update(ctx context.Context, post *domain.Post) (err error) {
	return nil
}

func (p *postUsecase) GetAll(ctx context.Context) (posts *[]domain.PostRequestDTO, err error) {
	return nil, nil
}

func (p *postUsecase) GetByID(ctx context.Context, id int64) (post *domain.PostRequestDTO, err error) {
	return nil, nil
}

func (p *postUsecase) Delete(ctx context.Context, id int64) (err error) {
	return nil
}
