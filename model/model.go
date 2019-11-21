package model

import (
	"time"
)

type BaseModel struct {
	Id       uint64     `gorm: "primary_key;AUTO_INCREMENT;column:id" json:"-"`
	CreateAt time.Time  `gorm: "column:createAt" json: "-"`
	UpdateAt time.Time  `gorm: "column:updateAt" json: "-"`
	DeleteAt *time.Time `gorm: "column:deleteAt" sql:"index" json: "-"`
}
