package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/zhe-ma/login-server-study/handler"
	"github.com/zhe-ma/login-server-study/model"

	"github.com/zhe-ma/login-server-study/pkg/constvar"
	"github.com/zhe-ma/login-server-study/pkg/errno"
)

func List(c *gin.Context) {
	param, _ := strconv.Atoi(c.Param("offset"))
	offset := uint64(param)

	param, _ = strconv.Atoi(c.Param("limit"))
	limit := uint64(param)
	if limit == 0 || limit > constvar.MaxLimit {
		limit = constvar.DefaultLimit
	}

	username := c.Param("username")

	log.Infof("List users. limit:%d, offset:%d, username:%s.", limit, offset, username)

	totalCount, userModels, err := model.ListUsers(username, limit, offset)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	userInfos := make([]UserInfo, 0)
	for _, userModel := range *userModels {
		append()
	}

	// userInfos := &ListResponse{
	// 	TotalCount: totalCount,
	// }

	handler.SendResponse(c, nil, userInfos)
}
