package web

import (
	"encoding/json"

	"github.com/Li-Khan/my-forum/delivery/web/middleware"
	"github.com/Li-Khan/my-forum/domain"
	"github.com/Li-Khan/my-forum/lib/mistake"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setSession(c *gin.Context, user *domain.User) error {
	b, err := json.Marshal(user)
	if err != nil {
		return mistake.ErrBadRequest
	}

	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 3600 * 24, Path: "/", HttpOnly: true})
	session.Set(middleware.Key, b)

	err = session.Save()
	if err != nil {
		return err
	}

	return nil
}
