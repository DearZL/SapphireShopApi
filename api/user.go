package api

import (
	"SapphireShop/SapphireShop_api/global"
	"SapphireShop/SapphireShop_api/model"
	"SapphireShop/SapphireShop_api/proto/user_srv"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func UserList(c *gin.Context) {
	req := &user_srv.PageInfo{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		global.Logger.Info("paramError!")
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("查询用户列表失败"))
		return
	}
	//连接rpc
	dial, err := grpc.Dial(fmt.Sprintf("%s:%d",
		global.Config.GetString("userServer.host"),
		global.Config.GetInt("userServer.port")),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Logger.Info("serverError!")
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("查询用户列表失败"))
		return
	}
	userClient := user_srv.NewUserSrvClient(dial)
	//调用rpc服务
	list, err := userClient.GetUserList(context.Background(), req)
	if err != nil {
		global.Logger.Info("查询用户列表失败")
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("查询用户列表失败"))
		return
	}
	c.JSON(http.StatusOK, model.Success("查询成功", list))
	return
}

func UserReg(c *gin.Context) {
	req := &user_srv.CreateUserInfo{}
	err := c.ShouldBindJSON(req)
	if err != nil {
		global.Logger.Info("paramError!")
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("用户注册失败"))
		return
	}
	//连接rpc
	dial, err := grpc.Dial(fmt.Sprintf("%s:%d",
		global.Config.GetString("userServer.host"),
		global.Config.GetInt("userServer.port")),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Logger.Info("serverError!")
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("用户注册失败"))
		return
	}
	userClient := user_srv.NewUserSrvClient(dial)
	res, err := userClient.CreateUser(context.Background(), req)
	if err != nil {
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("用户注册失败"))
		return
	}
	c.JSON(http.StatusOK, model.Success("用户注册成功", res))
	return
}
