package configs

import (
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
