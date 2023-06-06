package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnvVariables(){
	err:=godotenv.Load()

	if(err != nil){
		log.Fatal("Error loading env file")
	}
}
var DB *gorm.DB

func ConnectToDB() {
	dsn := "root:password@tcp(localhost:3306)/Movies?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get underlying *sql.DB")
	}
	fmt.Println("DB connected at: ",sqlDB)
}
