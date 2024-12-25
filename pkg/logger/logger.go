package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()

	// 设置日志格式为JSON
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	// 创建logs目录
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic(fmt.Sprintf("创建日志目录失败: %v", err))
	}

	// 设置日志文件
	filename := path.Join("logs", fmt.Sprintf("%s.log", time.Now().Format("2006-01-02")))
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("打开日志文件失败: %v", err))
	}

	// 同时输出到文件和控制台
	Log.SetOutput(file)
	Log.AddHook(NewConsoleHook())

	// 设置日志级别
	Log.SetLevel(logrus.InfoLevel)
}

// ConsoleHook 用于同时输出到控制台
type ConsoleHook struct{}

func NewConsoleHook() *ConsoleHook {
	return &ConsoleHook{}
}

func (hook *ConsoleHook) Fire(entry *logrus.Entry) error {
	str, err := entry.String()
	if err != nil {
		return err
	}
	fmt.Printf("%s", str)
	return nil
}

func (hook *ConsoleHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
