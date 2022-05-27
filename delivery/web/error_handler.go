package web

import (
	"log"
	"net/http"

	"github.com/Li-Khan/my-forum/lib/mistake"
	"github.com/gin-gonic/gin"
)

func errorHandler(c *gin.Context, status int, err error) {
	log.Println(err)
	c.JSON(status, err)
}

func getStatusCode(err error) (code int) {
	switch err {
	case mistake.ErrNotFound:
		return http.StatusNotFound
	case mistake.ErrBadRequest:
		return http.StatusBadRequest
	case mistake.ErrConflict:
		return http.StatusConflict
	case mistake.ErrUnauthorized:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
