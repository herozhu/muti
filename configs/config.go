package configs

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	//初始化日志文件
	if err := c.initLog(); err != nil {
		return err
	}

	// 监控配置文件变化并热加载程序
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")     // 设置配置文件格式为YAML
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) initLog() error {
	baseLogPath := path.Join(viper.GetString("app.LogSavePath"), viper.GetString("app.LogSaveName"))
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),      // 生成软链，指向最新日志文件(这里指向的是路径)
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
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
	}, &logrus.TextFormatter{
		TimestampFormat: viper.GetString("app.TimeFormat"), //设置时间格式
	})

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: viper.GetString("app.TimeFormat"), //设置时间格式
		FullTimestamp:   true,
		ForceColors:     true, //设置控制器输出颜色
	})

	//设置输出颜色
	logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	//设置最低loglevel
	logrus.SetLevel(logrus.InfoLevel)

	logrus.AddHook(lfHook)

	return nil

}

// 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}

//func InitConf() {
//	var err error
//	cfg, err = ini.Load("./configs/config.yaml")
//	if err != nil {
//		logrus.Fatalf("Fail to parse './config.yaml': %v", err)
//	}
//
//	mapTo("app", AppConfig)
//	mapTo("server", ServerConfig)
//	mapTo("database", DatabaseConfig)
//	mapTo("redis", RedisConfig)
//
//	AppConfig.ImageMaxSize = AppConfig.ImageMaxSize * 1024 * 1024
//	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
//	ServerConfig.WriteTimeout = ServerConfig.ReadTimeout * time.Second
//	RedisConfig.IdleTimeout = RedisConfig.IdleTimeout * time.Second
//}
//
//func mapTo(section string, v interface{}) {
//	err := cfg.Section(section).MapTo(v)
//	if err != nil {
//		logs.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
//	}
//}
