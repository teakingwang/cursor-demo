package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitRouter(logger *logrus.Logger) *gin.Engine {
	r := gin.Default()

	// API 路由组
	api := r.Group("/api")
	{
		// 在这里添加您的路由
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	logger.Info("路由初始化完成")
	return r
}
