package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"github.com/muti/models"
	"github.com/muti/pkg/errno"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

var (
	dao      = models.Tags{}
	validate *validator.Validate
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetTags(w http.ResponseWriter, r http.Request) {
	defer r.Body.Close()
	var tags []models.Tags
	tags, err := dao.FindAllTags()
	if err != nil {
		responseWithJson(w, http.StatusOK, errno.ERROR_GET_TAGS_FAIL)
		return
	}
	responseWithJson(w, http.StatusOK, tags)

}

func GetTag(c *gin.Context) {

}
