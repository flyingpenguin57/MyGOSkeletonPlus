package routerConfig

import (
	"github.com/gin-gonic/gin"
)

func CreateRouter(router *gin.Engine) {

	// 创建一个基础路径为 /api 的路由组
	apiGroup := router.Group("/api")
	{
		// 在 /api 路由组中创建一个 /users 路由组
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/register", )
			userGroup.POST("/login", )
			userGroup.POST("/userInfo", )
			userGroup.POST("/updateUserInfo", )
			userGroup.POST("/deleteUser", )
		}
	}
}
