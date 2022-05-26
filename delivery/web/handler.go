package web

import (
	"github.com/Li-Khan/my-forum/domain"
	"github.com/gin-gonic/gin"
)

// Handler...
type Handler struct {
	UserUsecase        domain.UserUsecase
	PostUsecase        domain.PostUsecase
	CommentUsecase     domain.CommentUsecase
	TagUsecase         domain.TagUsecase
	PostTagUsecase     domain.PostTagUsecase
	VotePostUsecase    domain.VotePostUsecase
	VoteCommentUsecase domain.VoteCommentUsecase
}

func NewHandler(r *gin.Engine, h *Handler) {
	// user := r.Group("/user")

	// post := r.Group("/post")

	// postVote := r.Group("/post/vote")

	// comment := r.Group("/comment")

	// commentVote := r.Group("/comment/vote")
}
