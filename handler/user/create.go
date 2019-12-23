package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/zhe-ma/login-server-study/handler"
	"github.com/zhe-ma/login-server-study/model"

	"github.com/zhe-ma/login-server-study/pkg/errno"
	"github.com/zhe-ma/login-server-study/util"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.BasicUserRequest true "Create a new user"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	log.Info("User create function called.", lager.Data{"X-Request-Id": util.GetRequestId(c)})

	var r BasicUserRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	if err := u.Validate(); err != nil {
		log.Error("Failed to validate user data.", err)
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		log.Error("Failed to encrypt user password.", err)
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	if err := u.Create(); err != nil {
		log.Error("Failed to create user in DB.", err)
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
