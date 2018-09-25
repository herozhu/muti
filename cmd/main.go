package main

import (
	"github.com/muti/configs"
	"github.com/muti/models"
	"github.com/muti/routers"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	//初始化配置
	pflag.Parse()
	if err := configs.Init(*cfg); err != nil {
		panic(err)
	}

	models.DB.Init()
	defer models.DB.Close()

	routers.InitRouter()

}
