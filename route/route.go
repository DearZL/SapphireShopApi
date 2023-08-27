package route

import (
	"SapphireShop/SapphireShop_api/api"
	"github.com/gin-gonic/gin"
)

func InitRouteGroup(r *gin.Engine) {
	baseUrl := "/api"
	route := r.Group(baseUrl)

	userAuth := route.Group("user")
	{
		userAuth.GET("list", api.UserList)
		userAuth.GET("userInfo")
	}
	userNoAuth := route.Group("user")
	{
		userNoAuth.POST("login")
		userNoAuth.POST("reg", api.UserReg)
	}

	commonAuth := route.Group("common")
	{
		commonAuth.GET("")
	}
	commonNoAuth := route.Group("common")
	{
		commonNoAuth.GET("sendCode/:email", api.SendCode)
	}

}
