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

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	log.Infof("Get user info. User id: %d.", id)

	u, err := model.GetUser(id)

	// 如何区分无此人，和数据库操作失败还是rowaffected?。
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, u)
}
