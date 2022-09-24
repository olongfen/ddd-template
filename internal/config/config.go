package config

import (
	"bytes"
	"ddd-template/pkg/utils"
	"github.com/fsnotify/fsnotify"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Configs struct {
	HTTP      HTTP
	Database  Database
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
	Driver      string
	Host        string
	Port        string
	User        string
	Password    string
	DBName      string
	TimeZone    string
	SSLMode     string
	Debug       bool
	AutoMigrate bool
}

type HTTP struct {
	Host string
	Port string
}

func Get() *Configs {
	if conf == nil {
		conf = new(Configs)
	}
	return conf
}

var conf *Configs

// setDefault 默认配置，文件改变会改变
func setDefault() {
	viper.SetDefault("http", HTTP{
		Host: "0.0.0.0",
		Port: "8818",
	})
	viper.SetDefault("languages", []string{"cn", "en"})
	viper.SetDefault("database", Database{
		Driver:      "postgresql",
		Host:        "127.0.0.1",
		Port:        "5432",
		User:        "postgres",
		Password:    "123456",
		DBName:      "postgres",
		TimeZone:    "Asia/Shanghai",
		SSLMode:     "disable",
		Debug:       true,
		AutoMigrate: true,
	})
	viper.SetDefault("log", Log{
		Filename:   "./log/server.log",
		ErrorFile:  "./log/server-err.log",
		MaxSize:    20,
		MaxBackups: 50,
		MaxAges:    30,
		Compress:   true,
		Debug:      true,
	})
}

// set 设置一些固定不可改变的配置
func set() {

}

func InitConfigs(confPath string) *Configs {
	var (
		err       error
		globalCfg = Get()
	)
	setDefault()
	// 载入配置
	viper.SetConfigFile(confPath)
	viper.SetConfigType("yaml")
	_, err = os.Stat(confPath)
	// 配置文件不存在自动读取默认配置然后创建创建
	if err != nil {
		if err = viper.WriteConfigAs(confPath); err != nil {
			log.Fatalln(err)
		}
		if err = viper.Unmarshal(globalCfg); err != nil {
			log.Fatal(err)
		}
	} else {
		var (
			originalBytes  []byte
			originalStruct = new(Configs)
			changeBytes    []byte
		)
		// 读取旧文件含有的配置
		if originalBytes, err = os.ReadFile(confPath); err != nil {
			log.Fatalln(err)
		}
		if err = yaml.Unmarshal(originalBytes, originalStruct); err != nil {
			log.Fatalln(err)
		}
		if err = utils.Copier(viper.AllSettings(), globalCfg); err != nil {
			log.Fatalln(err)
		}

		// 自动添加新的字段
		if err = utils.Copier(originalStruct, globalCfg); err != nil {
			log.Fatalln(err)
		}
		if changeBytes, err = jsoniter.Marshal(globalCfg); err != nil {
			log.Fatalln(err)
		}
		if err = viper.ReadConfig(bytes.NewReader(changeBytes)); err != nil {
			log.Fatalln(err)
		}
		if err = viper.WriteConfig(); err != nil {
			log.Fatalln(err)
		}

	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(globalCfg); err != nil {
			log.Fatal(err)
		}
	})

	return globalCfg
}
