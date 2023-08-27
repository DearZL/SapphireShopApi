package api

import (
	"SapphireShop/SapphireShop_api/common/utils"
	"SapphireShop/SapphireShop_api/global"
	"SapphireShop/SapphireShop_api/model"
	"SapphireShop/SapphireShop_api/proto/email_srv"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"strconv"
	"time"
)

func SendCode(c *gin.Context) {
	//连接rpc
	dial, err := grpc.Dial(fmt.Sprintf("%s:%d",
		global.Config.GetString("emailServer.host"),
		global.Config.GetInt("emailServer.port")),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Logger.Info("serverError!")
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("发送验证码失败"))
		return
	}
	emailSrv := email_srv.NewEmailSrvClient(dial)
	code := strconv.Itoa(utils.RandInt(111111, 999999))
	req := &email_srv.Email{
		Email:   c.Param("email"),
		Subject: "验证码",
		Msg:     "您的验证码为: " + code + "  五分钟内有效.",
		Code:    code,
		Expire:  int64(5 * time.Minute),
	}
	_, err = emailSrv.SendCode(context.Background(), req)
	if err != nil {
		global.Logger.Info(err.Error())
		c.JSON(http.StatusOK, model.Fail("验证码发送失败"))
		return
	}
	c.JSON(http.StatusOK, model.Success("发送验证码成功", nil))
}
