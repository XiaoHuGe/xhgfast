package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:64;not null"`
}

func (this *User) IsExist(telephone string) bool {

	if err := DB.Where("telephone= ?", telephone).First(this).Error; err == nil {
		return true
	} else if err != nil && err == gorm.ErrRecordNotFound {
		return false
	} else {
		//err = errors.New("数据库内部错误")
		panic(err)
		return false
	}
	return false
}

func (this *User) GetUserById(id int) error {

	if err := DB.Where("id= ?", id).First(this).Error; err != nil {
		return err
	}
	return nil
}
