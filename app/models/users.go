package models

import (
	"github.com/micro/go-micro/util/log"
	"github.com/spadesk1991/skip-micro/app/dao"
)

type Users struct {
	ID   int    `json:"id" gorm:"column:id;auto_increment;primary_key"`
	Name string `json:"name=" gorm:"column:name;varchar(23);unique"`
	PWD  string `json:"pwd" gorm:"column:pwd;varchar(20)"`
}

func init() {
	if e := dao.GetDB().AutoMigrate(Users{}).Error; e != nil {
		log.Error(e)
	}
}

func (s Users) TableName() string {
	return "user"
}
