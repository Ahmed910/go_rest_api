package entity

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type DBConnection struct {
	connection *gorm.DB
}

var db DBConnection

func NewConnectionToDB() {
	if db.connection != nil {
		return
	}
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("APP_LOCAL"), os.Getenv("MYSQL_PORT"), os.Getenv("DB_NAME"))
	dbLayer, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) // sqlite.Open("test.db")
	if err != nil {
		panic("Failed To Connect To DB")
	}
	dbLayer.AutoMigrate(&Author{}, &Video{})
	db.connection = dbLayer
}

func (dbConnection *DBConnection) CloseDB() {
	mysqlDB, err := dbConnection.connection.DB()

	if err != nil {
		panic("Failed To Close To DB")
	}
	mysqlDB.Close()
}
