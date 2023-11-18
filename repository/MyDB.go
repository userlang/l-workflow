package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义数据库连接配置
type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// 数据库连接全局变量
var DB *gorm.DB

// 初始化数据库连接
func InitDB() (*gorm.DB, error) {

	var config = Config{"gbjk@test", "gbjk@testpwd1", "172.16.0.104", "13049", "gbjk_workflow"}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
