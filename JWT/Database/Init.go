package Database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

const (
	UserName     string = "gary"
	Password     string = "867123e578n"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "test"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

var conn *gorm.DB

func Init() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Addr, Port,
		Database)
	conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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
		if err := Init(); err != nil {
			time.Sleep(5 * time.Second)
		}
	}
}
