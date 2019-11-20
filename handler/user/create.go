package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/zhe-ma/login-server-study/pkg/errno"
)

func Create(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var err error
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": errno.ErrBind})
		return
	}

	log.Debugf("username: %s, password: %s", user.Username, user.Password)

	if user.Username == "" {
		err := errno.New(errno.ErrUserNotFound, fmt.Errorf("username can't be empty"))
		log.Errorf(err, "Get an error")
	}

	if user.Password == "" {
		err = fmt.Errorf("pasword is empty")
	}

	code, message := errno.DecodeError(err)
	c.JSON(http.StatusOK, gin.H{"code": code, "message", message})
}
