package model

import (
	"Gorm/database"
)

type User struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

var conn = database.GetDB()

func (User) List(input User) (output []User, err error) {
	//fmt.Println(input.ID)
	var users []User
	tx := conn.Model(&User{}) //幫找資料的做好model
	if input.ID != 0 {
		tx = tx.Where("id = ?", input.ID)
	}
	if err = conn.Model(&User{}).Where("id = ?", input.ID).Find(&users).Error; err != nil {
		return
	}
	return users, err
}
func (User) Update(input User) (err error) {
	if err = conn.Model(&User{}).Where("id = ?", input.ID).Update("name", input.Name).Error; err != nil {
		return
	}
	return
}
func (User) Delete(input int64) (err error) {
	if err = conn.Model(&User{}).Delete(&User{}, input).Error; err != nil {
		return
	}
	return
}

// TableName :預設資料表會以struct的名字加上s為表名，如果struct名字與資料表不一樣可以使用TableName複寫
func (User) TableName() string {
	return "users"
}
