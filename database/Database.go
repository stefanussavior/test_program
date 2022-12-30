package database

import (
	"fmt"
	"test_netwerk/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	USER := "root"
	PASSWORD := ""
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "test_netwerk"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASSWORD, HOST, PORT, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Koneksi Berhasil")
	db.AutoMigrate(&models.Investasi{})
	return db
}
