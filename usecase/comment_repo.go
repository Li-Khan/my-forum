package usecase

import (
	"context"

	"github.com/Li-Khan/my-forum/domain"
)

type commentUsecase struct {
	commentRepo domain.CommentRepository
}

func NewCommentUsecase(c domain.CommentRepository) domain.CommentUsecase {
	return &commentUsecase{
		commentRepo: c,
	}
}

func (c *commentUsecase) Create(ctx context.Context, comment *domain.Comment) (id int64, err error) {
	return 0, nil
}

func (c *commentUsecase) Update(ctx context.Context, comment *domain.Comment) (err error) {
	return nil
}

func (c *commentUsecase) GetByPostID(ctx context.Context, id int64) (comments []*domain.CommentRequestDTO, err error) {
	return nil, nil
}

func (c *commentUsecase) GetByUserID(ctx context.Context, id int64) (comments []*domain.CommentRequestDTO, err error) {
	return nil, nil
}

func (c *commentUsecase) Delete(ctx context.Context, id int64) (err error) {
	return nil
}
