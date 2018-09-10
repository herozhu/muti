package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muti/api/v1"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	ApiV1 := r.Group("/api/v1")

	{
		ApiV1.GET("/user", v1.GetUser)
		ApiV1.GET("/tags", v1.GetTag)

	}

	return r

}
