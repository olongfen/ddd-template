package initialization

import (
	"ddd-template/internal/common/conf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitConf(confPath string) *conf.Configs {
	var (
		err error
		c   = conf.Get()
	)
	_, err = os.Stat(confPath)
	if err != nil {
		log.Fatalln(err)
	}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(confPath)
	_ = viper.ReadInConfig()
	if err = viper.Unmarshal(c); err != nil {
		log.Fatal(err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(c); err != nil {
			log.Fatal(err)
		}
	})
	return c
}
