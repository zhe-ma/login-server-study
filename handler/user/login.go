package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"

	"github.com/zhe-ma/login-server-study/handler"
	"github.com/zhe-ma/login-server-study/model"
	"github.com/zhe-ma/login-server-study/pkg/errno"
	"github.com/zhe-ma/login-server-study/pkg/token"
	"github.com/zhe-ma/login-server-study/util"
)

// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
func Login(c *gin.Context) {
	var r BasicUserRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	log.Debugf("Username: %s, password: %s.", r.Username, r.Password)

	u, err := model.GetUserByName(r.Username)
	// 如何区分无此人，和数据库操作失败还是rowaffected?。
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "Failed to get user [%s].", r.Username)
		return
	}

	if err := util.ComparePassword(u.Password, r.Password); err != nil {
		handler.SendResponse(c, errno.ErrInvalidPassword, nil)
		return
	}

	t, err := token.Sign(token.Context{ID: u.ID, Username: u.Username})
	if err != nil {
		handler.SendResponse(c, errno.ErrSignToken, nil)
		return
	}

	handler.SendResponse(c, nil, Token{Token: t})
}
