package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/muti/api/v1"
)

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	t := r.Group("/api/v1/tags")

	{
		t.GET("", v1.GetTags)

	}

	a := r.Group("/api/v1/articles")
	{
		a.GET("")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
	return r
}
