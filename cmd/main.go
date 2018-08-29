package main

import (
	"github.com/muti/configs"
	"github.com/muti/models"
	"github.com/sirupsen/logrus"
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

	models.InitData()
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	logrus.Fatal("Bye.")         //log之后会调用os.Exit(1)
	logrus.Panic("I'm bailing.") //log之后会panic()
}
