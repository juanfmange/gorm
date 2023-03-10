package main

import (
	"github.com/jinzhu/gorm"
	"github.com/juanfmange/gorm/libs"
)

type User struct {
	gorm.Model
	//ID       uint    `gorm:"primary_key;auto_increment"`
	Account  string `gorm:"type:varchar(20);not null;index:username"`
	Password string `gorm:"type:char(32);not null;"`
}

func main() {
	dbConfig := libs.Configure("./", "mysql")
	libs.DB = dbConfig.InitMysqlDB()
	libs.DB.AutoMigrate(&User{})

}
