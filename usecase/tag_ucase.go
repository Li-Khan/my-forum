package usecase

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
)

type tagUsecase struct {
	tagRepo domain.TagRepository
}

func NewTagUsecase(t domain.TagRepository) domain.TagUsecase {
	return &tagUsecase{
		tagRepo: t,
	}
}

func (t *tagUsecase) Create(ctx context.Context, tag *domain.Tag) (id int64, err error) {
	return 0, nil
}
