package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"workflow/config"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySqlUsername, config.MySqlPassword, config.MySqlHost, config.MySqlPort, config.MySqlDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}

func Start() *gorm.DB {
	tx := DB.Begin()
	// 如果发生错误，则回滚事务并返回错误

	if tx.Error != nil {
		tx.Rollback()
	}
	return tx
}

func End(db *gorm.DB) {
	db.Commit()
}
