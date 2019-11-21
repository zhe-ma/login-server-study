package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhe-ma/login-server-study/pkg/errno"
)

type Response struct {
	Code    int         `json: "code"`
	Message string      `json: "message"`
	Data    interface{} `json: "data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeError(err)
	c.JSON(http.StatusOK, Response{
		code:    code,
		Message: message,
		Data:    data,
	})
}
