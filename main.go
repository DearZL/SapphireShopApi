package main

import (
	"SapphireShop/SapphireShop_api/common/config"
	"SapphireShop/SapphireShop_api/global"
	"SapphireShop/SapphireShop_api/route"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	global.InitConfig()
	global.InitLogConfig()
}

func main() {
	defer config.SyncLog(global.Logger)
	ip := global.Config.GetString("app.host")
	port := global.Config.GetInt("app.port")

	r := gin.Default()
	gin.SetMode(global.Config.GetString("app.mode"))
	_ = r.SetTrustedProxies(nil)

	route.InitRouteGroup(r)

	global.Logger.Info("RouteGroup注册完成")
	err := r.Run(fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic(err)
		return
	}
}
