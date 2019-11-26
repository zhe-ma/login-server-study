package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/zhe-ma/login-server-study/handler"
	"github.com/zhe-ma/login-server-study/model"

	"github.com/zhe-ma/login-server-study/pkg/errno"
)

func Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	log.Infof("Get user info. User id: %d.", id)

	u, err := model.GetUser(uint64(id))

	// 如何区分无此人，和数据库操作失败还是rowaffected?。
	if err != nil {
		log.Info(err.Error())
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	handler.SendResponse(c, nil, u)
}
