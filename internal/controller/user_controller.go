package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/cursor-demo/internal/service"
	"github.com/teakingwang/cursor-demo/pkg/logger"
	"github.com/teakingwang/cursor-demo/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (c *UserController) List(ctx *gin.Context) {
	logger.Log.Info("获取用户列表")
	users, err := c.userService.GetUsers()
	if err != nil {
		logger.Log.Errorf("获取用户列表失败: %v", err)
		ctx.JSON(500, response.Error(500, "获取用户列表失败"))
		return
	}
	ctx.JSON(200, response.Success(users))
}

func (c *UserController) Get(ctx *gin.Context) {
	// 实现获取单个用户的逻辑
}

func (c *UserController) Create(ctx *gin.Context) {
	// 实现创建用户的逻辑
}

func (c *UserController) Update(ctx *gin.Context) {
	// 实现更新用户的逻辑
}

func (c *UserController) Delete(ctx *gin.Context) {
	// 实现删除用户的逻辑
}
