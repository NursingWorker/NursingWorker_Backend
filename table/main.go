package main

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"nursing_work/config"
	"nursing_work/model"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
	)
	var err error
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := config.Init("./conf/config.yaml", ""); err != nil {
		log.Println(err)
	}
	InitDB()
	DB.AutoMigrate(&model.Comment{}, &model.User{}, &model.Token{}, &model.Post{},
		&model.Collection{}, &model.PostFl{}, &model.View{}, &model.Videos{},
		&model.Follow{}, &model.Like{}, &model.Reply{}, &model.Images{})
}
