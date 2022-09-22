package config

import (
	"ddd-template/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Configs struct {
	HTTP      HTTP
	Database  Database
	Debug     bool
	Language  string
	Languages []string
	Log       Log
}

type Log struct {
	Filename   string
	ErrorFile  string
	MaxSize    int
	MaxBackups int
	MaxAges    int
	Compress   bool
	Debug      bool
}

type Database struct {
	Driver string
	Source string
	Dev    bool
}

type HTTP struct {
	Host string
	Port string
}

func Get() *Configs {
	return Confs
}

var Confs = new(Configs)

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

func InitConfigs(confPath string) *Configs {
	var (
		err error
		_   os.FileInfo
		c   = Get()
	)
	viper.SetConfigType("yaml")
	_, err = os.Stat(confPath)
	// 配置文件不存在自动创建
	if err != nil {
		var (
			file *os.File
			b    []byte
		)
		newCfg := new(Configs)
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
			fileCfg    = new(Configs)
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
