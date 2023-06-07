package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}
}

var DB *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	dsn := "root:password@tcp(localhost:3306)/Movies?charset=utf8mb4&parseTime=True&loc=Local"
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
func Close() {

}
