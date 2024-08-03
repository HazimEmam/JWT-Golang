package database

import (
	"fmt"

	"github.com/HazimEmam/JWTtutorial/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var DB *gorm.DB

func InitDB( cfg Config){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	db, err:= gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if err != nil{
		fmt.Println("Error in Connection")
		panic(err)
	}

	//Migration
	if err := db.AutoMigrate(&models.User{}); err != nil{
		fmt.Println("Error in Migration")
		panic(err)
	}

	fmt.Println("Migrated database")

	DB = db
}