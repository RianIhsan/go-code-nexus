package database

import (
	"fmt"
	"github.com/RianIhsan/go-code-nexus/config"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func BootDatabase(cnf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cnf.Database.DbHost, cnf.Database.DbUser, cnf.Database.DbPass, cnf.Database.DbName, cnf.Database.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open database", err.Error())
		return nil
	}
	log.Info("Database connected")
	return db
}
