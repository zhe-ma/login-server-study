package model

import (
	"fmt"

	"github.com/lexkong/log"
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

func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *UserModel) Encrypt() (err error) {
	u.Password, err = util.Encrypt(u.Password)
	return err
}

func (u *UserModel) Create() error {
	return DB.Self.Create(&u).Error
}

func GetUser(id uint64) (*UserModel, error) {
	user := &UserModel{}
	db := DB.Self.Where("id = ?", id).First(&user)
	return user, db.Error
}

func GetUserByName(username string) (*UserModel, error) {
	user := &UserModel{}
	db := DB.Self.Where("username = ?", username).First(&user)
	return user, db.Error
}

func DeleteUser(id uint64) error {
	u := &UserModel{}
	u.BaseModel.ID = id
	return DB.Self.Delete(&u).Error
}

func (u *UserModel) Update() error {
	return DB.Self.Save(&u).Error
}

func ListUsers(username string, limit uint64, offset uint64) (uint64, []*UserModel, error) {
	var totalCount uint64 = 0
	userInfos := make([]*UserModel, 0)

	where := fmt.Sprintf("username like '%%%s%%'", username)
	log.Debugf("QueryUsers where sql:%s.", where)

	if err := DB.Self.Model(&UserModel{}).Where(where).Count(&totalCount).Error; err != nil {
		return totalCount, userInfos, err
	}

	if err := DB.Self.Where(where).Limit(limit).Offset(offset).Order("id ASC").Find(&userInfos).Error; err != nil {
		return totalCount, userInfos, err
	}

	return totalCount, userInfos, nil
}
