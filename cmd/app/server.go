package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/teakingwang/cursor-demo/config"
	"github.com/teakingwang/cursor-demo/pkg/redis"
)

type Server struct {
	port   string
	engine *gin.Engine
	logger *logrus.Logger
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		logger: logrus.New(),
	}
}

func (s *Server) Start() error {
	if s.logger == nil {
		s.logger = logrus.New()
	}

	// 设置 gin 的模式
	gin.SetMode(gin.ReleaseMode)

	s.logger.Info("Starting server on port ", s.port)
	// 初始化配置
	config.InitConfig()
	s.logger.Info("配置初始化完成")

	// 初始化数据库
	if err := config.InitDB(); err != nil {
		return fmt.Errorf("初始化数据库失败: %v", err)
	}
	s.logger.Info("数据库初始化完成")

	// 初始化Redis
	if err := redis.InitRedis(); err != nil {
		return fmt.Errorf("初始化Redis失败: %v", err)
	}
	s.logger.Info("Redis初始化完成")

	// 初始化路由
	s.engine = InitRouter(s.logger)
	s.logger.Info("路由初始化完成")

	// 优雅关机
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		s.logger.Info("正在关闭服务器...")
		os.Exit(0)
	}()

	s.logger.Infof("服务器启动在端口 %s", s.port)
	return s.engine.Run(fmt.Sprintf(":%s", s.port))
}
