package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/zhe-ma/login-server-study/handler/handler"
	"github.com/zhe-ma/login-server-study/model"
	"github.com/zhe-ma/login-server-study/pkg/errno"
	"github.com/zhe-ma/login-server-study/util"
)

func Create(c *gin.Context) {
	log.Info("User create function called.", lager.Data{"X-Request-Id": util.GetRequestId(c)})

	var r CeateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
