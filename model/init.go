package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%v:%v@tcp(43.138.61.49:3306)/nursing?charset=utf8mb4&parseTime=True&loc=Local", viper.GetString("db.username"), viper.GetString("db.password"))
	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}
	return nil
}
