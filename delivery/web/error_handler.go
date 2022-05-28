package web

import (
	"log"
	"net/http"

	"github.com/Li-Khan/my-forum/lib/mistake"
	"github.com/gin-gonic/gin"
)

type e struct {
	Error string `json:"error,omitempty"`
}

func errorHandler(c *gin.Context, status int, err error) {
	log.Println(err)
	c.IndentedJSON(status, e{Error: err.Error()})
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
