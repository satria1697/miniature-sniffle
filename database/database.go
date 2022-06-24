package database

import (
	"fmt"
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
	return db, nil
}
