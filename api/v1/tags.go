package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"github.com/muti/models"
	"github.com/muti/pkg/errno"
	"github.com/muti/tools"
	"github.com/spf13/viper"
	"net/http"
)

func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := errno.SUCCESS

	data["lists"] = models.GetTags(tools.GetPage(c), viper.GetInt("app.pagesize"), maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errno.GetMsg(code),
		"data": data,
	})
}
