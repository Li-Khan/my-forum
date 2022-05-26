package usecase

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
)

type postTagUsecase struct {
	postTagRepo domain.PostTagRepository
}

func NewPostTagUsecase(p domain.PostTagRepository) domain.PostTagUsecase {
	return &postTagUsecase{
		postTagRepo: p,
	}
}

func (p *postTagUsecase) Create(ctx context.Context, pt *domain.PostTag) (id int64, err error) {
	return 0, nil
}
