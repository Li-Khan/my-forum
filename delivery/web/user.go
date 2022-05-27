package web

import (
	"net/http"
	"strings"
	"time"

	"github.com/Li-Khan/my-forum/domain"
	"github.com/Li-Khan/my-forum/lib/mistake"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) signup(c *gin.Context) {
	user := domain.User{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := c.BindJSON(&user)
	if err != nil {
		errorHandler(c, http.StatusBadRequest, err)
		return
	}

	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
		errorHandler(c, http.StatusBadRequest, mistake.ErrBadRequest)
		return
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		errorHandler(c, http.StatusInternalServerError, err)
		return
	}

	user.UserID, err = h.UserUsecase.Create(c, &user)
	if err != nil {
		errorHandler(c, getStatusCode(err), err)
		return
	}

	user.Password = ""
	c.JSON(http.StatusCreated, user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
