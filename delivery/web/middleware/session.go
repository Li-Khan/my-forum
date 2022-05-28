package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const Key string = "forum"

func InSessionAccessControl(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		value := session.Get(Key)
		if value != nil {
			c.JSON(http.StatusUnprocessableEntity, nil)
			return
		}
		h(c)
	}
}

func NotSessionAccessControl(h gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		value := session.Get(Key)
		if value == nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		user := domain.User{}
		err := json.Unmarshal(value.([]byte), &user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		c.JSON(http.StatusUnauthorized, user)
		h(c)
	}
}
