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
	param, _ := strconv.Atoi(c.Query("offset"))
	offset := uint64(param)

	param, _ = strconv.Atoi(c.Query("limit"))
	limit := uint64(param)
	if limit == 0 || limit > constvar.MaxLimit {
		limit = constvar.DefaultLimit
	}

	username := c.Query("username")

	log.Infof("List users. limit:%d, offset:%d, username:%s.", limit, offset, username)

	totalCount, userModels, err := model.ListUsers(username, limit, offset)
	if err != nil {
		log.Error("Failed to queryUsers.", err)
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	userInfos := make([]*Info, 0)
	for _, userModel := range userModels {
		userInfos = append(userInfos, &Info{
			ID:       userModel.ID,
			Username: userModel.Username,
			Password: userModel.Password,
			CreateAt: userModel.CreateAt,
			UpdateAt: userModel.UpdateAt,
		})
	}

	listResponse := &ListResponse{
		TotalCount: totalCount,
		UserInfos:  userInfos,
	}

	handler.SendResponse(c, nil, listResponse)
}
