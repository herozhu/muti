package tools

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * viper.GetInt("app.PageSize")
	}

	return result
}
