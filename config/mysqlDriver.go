package config

import (
	"api/mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// database connection
func InitDB() {
	//declare struct configm& variable connectionString
	connectionString := "root:nephilim17@tcp(127.0.0.1:3306)/gorm_mvc?charset=utf8&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

// proses database migration
// Auto Update database
func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
