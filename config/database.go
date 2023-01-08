package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/riwandylbs/go-learning-gin-gonic/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetEnvValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load file env")
	}

	return os.Getenv(key)
}

func dsn() string {
	host := GetEnvValue("DB_HOST")
	port := GetEnvValue("DB_PORT")
	db := GetEnvValue("DB_NAME")
	username := GetEnvValue("DB_USERNAME")
	pwd := GetEnvValue("DB_PASSWORD")

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, pwd, host, port, db)
}

var DB *gorm.DB

func DatabaseConnections() *gorm.DB {

	db, err := gorm.Open(mysql.Open(dsn()))
	if err != nil {
		panic(err)
	}

	// migrations
	var user []models.User
	db.AutoMigrate(&user)

	// setting database pooling connections
	sqlDB, err := db.DB() // more details informations visit https://gorm.io/docs/generic_interface.html

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(50)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	// jangan lupa di close connectionnya di function yang panggil
	// please use this command to close sqlDB.Close()
	return db
}
