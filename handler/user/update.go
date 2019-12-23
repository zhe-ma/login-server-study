package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/zhe-ma/login-server-study/handler"
	"github.com/zhe-ma/login-server-study/model"

	"github.com/zhe-ma/login-server-study/pkg/errno"
	"github.com/zhe-ma/login-server-study/util"
)

// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.UserModel true "The user info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /user/{id} [put]
func Update(c *gin.Context) {
	log.Info("User update function called.", lager.Data{"X-Request-Id": util.GetRequestId(c)})

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	u.ID = uint64(id)

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

	// 如何区分无此人，和数据库操作失败还是rowaffected?。
	if err := u.Update(); err != nil {
		log.Error("Failed to create user in DB.", err)
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
