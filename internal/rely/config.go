package rely

import (
	"bytes"
	"github.com/fsnotify/fsnotify"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

type Configs struct {
	HTTP        HTTP
	RPC         RPC
	Database    Database
	Log         Log
	Redis       Redis
	Nacos       Nacos
	EnableGraph bool
}

type Nacos struct {
	Register   bool   // true 注册网关
	ClientName string // 客户端地址
	Host       string // 网关host
	Port       uint64 // 网关port
	Weight     float64
}
type RPC struct {
	Host string
	Port uint
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
	Host    string
	Port    string
	BaseURL string
}

type Redis struct {
	Use       bool
	Addr      string
	Password  string
	DB        int
	KeyPrefix string
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
	viper.SetDefault("watchconfig", false)
	viper.SetDefault("http", HTTP{
		Host: "0.0.0.0",
		Port: "8818",
	})

	viper.SetDefault("database", Database{
		Driver:      "postgresql",
		Host:        "127.0.0.1",
		Port:        "5432",
		User:        "postgres",
		Password:    "business",
		DBName:      "business",
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
	viper.SetDefault("redis", Redis{
		Addr:      "127.0.0.1:6379",
		Password:  "123456",
		DB:        1,
		Use:       true,
		KeyPrefix: "system_manage",
	})

	viper.SetDefault("rpc", RPC{
		Host: "0.0.0.0",
		Port: 9060,
	})

	viper.SetDefault("nacos", Nacos{
		Register:   false,
		ClientName: "127.0.0.1",
		Host:       "127.0.0.1",
		Port:       10948,
		Weight:     0,
	})
	viper.SetDefault("enablegraph", false)
}

// set 设置一些固定不可改变的配置
func set() {

}

func InitConfigs(confPath string) (cfg *Configs, err error) {
	var (
		globalCfg = Get()
		base, _   = path.Split(confPath)
	)
	setDefault()
	if _, _err := os.Stat(base); _err != nil {
		_ = os.MkdirAll(base, os.ModePerm)
	}
	// 载入配置
	viper.SetConfigFile(confPath)
	viper.SetConfigType("yaml")
	_, err = os.Stat(confPath)
	// 配置文件不存在自动读取默认配置然后创建创建
	if err != nil {
		if err = viper.WriteConfigAs(confPath); err != nil {
			return
		}
		if err = viper.Unmarshal(globalCfg); err != nil {
			return
		}
	} else {
		var (
			changeBytes []byte
		)
		if err = viper.ReadInConfig(); err != nil {
			return
		}
		if err = viper.Unmarshal(globalCfg); err != nil {
			return
		}
		if changeBytes, err = jsoniter.Marshal(globalCfg); err != nil {
			err = errors.WithMessage(err, "Marshal")
			return
		}
		if err = viper.ReadConfig(bytes.NewBuffer(changeBytes)); err != nil {
			err = errors.WithMessage(err, "ReadConfig from changeBytes")
			return
		}
		if err = viper.WriteConfig(); err != nil {
			err = errors.WithMessage(err, "WriteConfig")
			return
		}

	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
		if err = viper.Unmarshal(globalCfg); err != nil {
			log.Fatal(err)
		}
	})

	log.Println("config init success")
	return globalCfg, nil
}
