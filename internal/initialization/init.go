package initialization

import (
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/xlog"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
)

func InitLog(cfg *conf.Configs) *zap.Logger {
	var (
		logger *zap.Logger
	)
	if cfg.Environment == "dev" {
		logger = xlog.NewDevelopment()
	} else {
		logger = xlog.NewProduceLogger()
	}
	xlog.Log = logger
	return logger
}

func InitConf(confPath string) *conf.Configs {
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
	if err = viper.Unmarshal(conf.Confs); err != nil {
		log.Fatal(err.Error())
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(conf.Confs); err != nil {
			log.Fatal(err)
		}
	})
	return conf.Confs
}
