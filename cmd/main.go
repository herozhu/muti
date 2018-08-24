package main

import (
	"github.com/muti/configs"
	"github.com/muti/models"
	"github.com/muti/pkg/logging"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.InitConf()
	logging.InitLog()
	models.InitData()
	logrus.Debug("Useful debugging information.")
	logrus.Info("Something noteworthy happened!")
	logrus.Warn("You should probably take a look at this.")
	logrus.Error("Something failed but I'm not quitting.")
	logrus.Fatal("Bye.")         //log之后会调用os.Exit(1)
	logrus.Panic("I'm bailing.") //log之后会panic()

}
