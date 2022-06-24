package database

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"six/user/domain"
	util "six/utils"
)

func InitDatabase(config util.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&domain.User{})
	initUser(db)

	return db, nil
}

func initUser(db *gorm.DB) {
	resDb := db.Find(&domain.User{})
	if resDb.RowsAffected > 0 {
		return
	}
	password, _ := bcrypt.GenerateFromPassword([]byte("inibebek"), bcrypt.DefaultCost)

	db.Create(&domain.User{
		Username: "inibebek",
		Email:    "bebek.com",
		Password: string(password),
	})
}
