package app

import (
	"github.com/gin-gonic/gin"
	"github.com/muti/pkg/errno"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  errno.GetMsg(errCode),
		"data": data,
	})

	return
}
