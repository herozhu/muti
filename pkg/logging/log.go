package logging

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/muti/configs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

const (
	MaxAge       = 7 * 24 * time.Hour
	RotationTime = 24 * time.Hour
)

func InitLog() {
	baseLogPath := path.Join(configs.AppConfig.LogSavePath, configs.AppConfig.LogSaveName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件(这里指向的是路径)
		rotatelogs.WithMaxAge(MaxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(RotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.TextFormatter{})

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: configs.AppConfig.TimeFormat, //设置时间格式
		FullTimestamp:   true,
		ForceColors:     true, //设置控制器输出颜色
	})

	//设置输出颜色
	logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)

	logrus.AddHook(lfHook)

}
