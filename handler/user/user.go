package user

import (
	"time"
)

type CeateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Info struct {
	ID       uint64    `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	CreateAt time.Time `json:"createTime"`
	UpdateAt time.Time `json:"updateTime"`
}

type ListResponse struct {
	TotalCount uint64  `json:"totalCount"`
	UserInfos  []*Info `json:"userInfos"`
}
