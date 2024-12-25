package database

import (
	"fmt"

	"github.com/teakingwang/cursor-demo/config"
	"github.com/teakingwang/cursor-demo/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.Database.Username,
		config.GlobalConfig.Database.Password,
		config.GlobalConfig.Database.Host,
		config.GlobalConfig.Database.Port,
		config.GlobalConfig.Database.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 自动迁移
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}

	DB = db
	return nil
}
