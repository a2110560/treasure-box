package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

const (
	UserName     string = ""
	Password     string = ""
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

var conn *gorm.DB

func connect() error {
	var err error
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	//連接MySQL
	conn, err = gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return err
	}
	//設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	_db, err := conn.DB()
	if err != nil {
		fmt.Println("get db failed:", err)
		return err
	}
	_db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	_db.SetMaxIdleConns(MaxIdleConns)
	_db.SetMaxOpenConns(MaxOpenConns)
	return err
}
func GetDB() *gorm.DB {
	for {
		if conn != nil {
			return conn
		}
		if err := connect(); err != nil {
			time.Sleep(5 * time.Second)
		}
	}
}
