package web

import (
	"github.com/Li-Khan/my-forum/delivery/web/middleware"
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

// NewHandler ...
func NewHandler(r *gin.Engine, h *Handler) {
	user := r.Group("/user")
	{
		user.POST("/signup", middleware.InSessionAccessControl(h.signup))
		user.POST("/signin", middleware.InSessionAccessControl(h.signin))
		user.POST("/signout", middleware.NotSessionAccessControl(h.signout))
	}

	// post := r.Group("/post")

	// postVote := r.Group("/post/vote")

	// comment := r.Group("/comment")

	// commentVote := r.Group("/comment/vote")
}
