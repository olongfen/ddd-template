package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Configs struct {
	Server      Server
	Database    Database
	JaegerHost  string
	Environment string
	ServiceId   int32
	ServiceName string
	Language    string
	Log         Log
}

type Log struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAges    int
	Compress   bool
}

type Database struct {
	Driver string
	Source string
}

type HTTP struct {
	Addr string
	Host string
	Port int
}

type Server struct {
	Http HTTP
}

func InitConf(confPath string) {
	var (
		err error
	)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(confPath)
	_ = viper.ReadInConfig()
	if err = viper.Unmarshal(conf); err != nil {
		zap.L().Sugar().Fatal(err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Sugar().Infof("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(conf); err != nil {
			zap.L().Sugar().Fatal(err)
		}
	})
}

func Get() Configs {
	return *conf
}

var conf = new(Configs)
