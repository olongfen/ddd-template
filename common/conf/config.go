package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Configs struct {
	Server      Server
	Database    Database
	JaegerHost  string
	Environment string
	Debug       bool
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

type GRpc struct {
	Host string
	Port int
}

type Server struct {
	Http HTTP
	GRpc GRpc
}

func InitConf(confPath string) {
	var (
		err error
	)
	_, err = os.Stat(confPath)
	if err != nil {
		log.Fatalln(err)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(confPath)
	_ = viper.ReadInConfig()
	if err = viper.Unmarshal(conf); err != nil {
		log.Fatal(err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(conf); err != nil {
			log.Fatal(err)
		}
	})
}

func Get() Configs {
	return *conf
}

var conf = new(Configs)
