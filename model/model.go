package model

import (
	"time"
)

type BaseModel struct {
	Id       uint64     `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreateAt time.Time  `gorm:"column:createdAt" json:"-"`
	UpdateAt time.Time  `gorm:"column:updatedAt" json:"-"`
	DeleteAt *time.Time `gorm:"column:deletedAt" sql:"index" json:"-"`
}
