package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/muti/models"
	"github.com/muti/pkg/app"
	"github.com/muti/pkg/errno"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var (
	dao = models.Tags{}
)

func GetTags(c *gin.Context) (r http.Response) {
	defer r.Body.Close()
	appG := app.Gin{C: c}
	var tags []models.Tags
	tags, err := dao.FindAllTags()
	if err != nil {
		appG.Response(http.StatusOK, errno.ERROR_GET_TAGS_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, errno.SUCCESS, tags)
	return
}

func GetTag(c *gin.Context) {
	id := bson.ObjectId.String("_id")
	appG := app.Gin{C: c}
	result, err := dao.FindTagById(id)
	if err != nil {
		appG.Response(http.StatusOK, errno.ERROR_NOT_EXIST_TAG, nil)
	}
	appG.Response(http.StatusOK, errno.SUCCESS, result)
}
