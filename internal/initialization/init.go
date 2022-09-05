package initialization

import (
	"ddd-template/internal/common/conf"
	"ddd-template/internal/common/utils"
	"ddd-template/internal/common/xlog"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func writeToFile(fileName string, content []byte) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	} else {
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt(content, n)
		f.Close()
	}
}

func InitConf(confPath string) *conf.Configs {
	var (
		err error
		_   os.FileInfo
		c   = conf.Get()
	)
	viper.SetConfigType("yaml")
	_, err = os.Stat(confPath)
	// 配置文件不存在自动创建
	if err != nil {
		var (
			file *os.File
			b    []byte
		)
		newCfg := new(conf.Configs)
		_ = viper.Unmarshal(newCfg)
		b, err = yaml.Marshal(newCfg)
		if file, err = os.Create(confPath); err != nil {
			log.Fatalln(err)
		}
		if _, err = file.Write(b); err != nil {
			log.Fatalln(err)
		}
		file.Close()
	} else {
		var (
			fileBytes  []byte
			viperBytes []byte
			fileCfg    = new(conf.Configs)
		)
		// 读取旧文件含有的配置
		if fileBytes, err = os.ReadFile(confPath); err != nil {
			log.Fatalln(err)
		}
		if err = yaml.Unmarshal(fileBytes, fileCfg); err != nil {
			log.Fatalln(err)
		}
		if err = mapstructure.Decode(viper.AllSettings(), c); err != nil {
			log.Fatalln(err)
		}

		// 自动添加新的字段
		if err = utils.Copier(fileCfg, c); err != nil {
			log.Fatalln(err)
		}
		if viperBytes, err = yaml.Marshal(c); err != nil {
			log.Fatalln(err)
		}
		writeToFile(confPath, viperBytes)

	}

	// 重新载入配置
	viper.SetConfigFile(confPath)
	if err = viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
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

func InitLog(cfg *conf.Configs) *zap.Logger {
	var (
		logger *zap.Logger
	)
	if cfg.Log.Debug {
		logger = xlog.NewDevelopment()
	} else {
		logger = xlog.NewProduceLogger()
	}
	xlog.Log = logger
	return logger
}

func InitI18N(cfg *conf.Configs) *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	if len(cfg.Languages) == 0 {
		panic("languages must be definition")
	}
	for _, v := range cfg.Languages {
		bundle.MustLoadMessageFile(fmt.Sprintf("internal/common/xi18n/active.%s.json", v))
	}
	return bundle
}
