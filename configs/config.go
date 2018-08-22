package configs

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string `ini:"JwtSecret"`
	PageSize  int    `ini:"pageSize"`
	PrefixUrl string `ini:"PrefixUrl"`

	RuntimeRootPath string `ini:"RuntimeRootPath"`

	ImageSavePath  string   `ini:"ImageSavePath"`
	ImageMaxSize   int      `ini:"ImageMaxSize"`
	ImageAllowExts []string `ini:"ImageAllowExts"`

	ExportSavePath string `ini:"ExportSavePath"`
	QrCodeSavePath string `ini:"QrCodeSavePath"`
	FontSavePath   string `ini:"FontSavePath"`

	LogSavePath string `ini:"LogSavePath"`
	LogSaveName string `ini:"LogSaveName"`
	LogFileExt  string `ini:"LogFileExt"`
	TimeFormat  string `ini:"TimeFormat"`
}

var AppConfig = &App{}

type Server struct {
	RunMode      string `ini:"RunMode"`
	HttpPort     int    `ini:"HttpPort"`
	ReadTimeout  string `ini:"ReadTimeout"`
	WriteTimeout string `ini:"WriteTimeout"`
}

var ServerConfig = &Server{}

type Database struct {
	Type        string `ini:"Type"`
	User        string `ini:"User"`
	PassWord    string `ini:"PassWord"`
	Host        string `ini:"Host"`
	Name        string `ini:"Name"`
	TablePrefix string `ini:"TablePrefix"`
}

var DatabaseConfig = &Database{}

type Redis struct {
	Host        string        `ini:"Host"`
	Password    string        `ini:"Password"`
	MaxIdle     string        `ini:"MaxIdle"`
	MaxActive   string        `ini:"MaxActive"`
	IdleTimeout time.Duration `ini:"IdleTimeout"`
}

var RedisConfig = &Redis{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("configs/config.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'configs/config.ini': %v", err)
	}

	mapTo("app", AppConfig)
	mapTo("server", ServerConfig)
	mapTo("database", DatabaseConfig)
	mapTo("redis", RedisConfig)

	AppConfig.ImageMaxSize = AppConfig.ImageMaxSize * 1024 * 1024
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.ReadTimeout * time.Second
	RedisConfig.IdleTimeout = RedisConfig.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
