package model

import (
	"github.com/zhe-ma/login-server-study/util"
	validator "gopkg.in/go-playground/validator.v9"
)

type UserModel struct {
	BaseModel
	Username string `gorm:"column:username;not null" binding:"required" validate:"min=1,max=32" json:"username"`
	Password string `gorm:"column:password;not null" binding:"required" validate:"min=5,max=128" json:"password"`
}

func (u *UserModel) TableName() string {
	return "tb_users"
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
func (u *UserModel) Encrypt() error {
	u.Password, err = util.Encrypt(u.Password)
	return err
}
