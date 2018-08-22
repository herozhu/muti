package logrus

import (
	"fmt"
	"github.com/muti/configs"
	"time"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", configs.AppConfig.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s", configs.AppConfig.LogSaveName, time.Now().Format(configs.AppConfig.TimeFormat), configs.AppConfig.LogFileExt)
}
