package db

import (
	"fmt"
	config "learngo/restapiserver/pkg/movies/configs"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ReturnDB() (*gorm.DB, error) {
	config, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBLink, config.DBPort, config.DBName)
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying *sql.DB: %w", err)
	}

	fmt.Println("DB connected at:", sqlDB)
	return DB, nil
}
