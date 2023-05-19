package rely

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
	"strings"
)

type Configs struct {
	HTTP        HTTP
	RPC         RPC
	DB          DB
	Log         Log
	Redis       Redis
	Nacos       Nacos
	EnableGraph bool
}

type Nacos struct {
	Register bool   // true 注册网关
	ClientIP string // 客户端地址
	Host     string // 网关host
	Port     uint64 // 网关port
	Weight   float64
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

type DB struct {
	Driver      string
	DSN         string
	Debug       bool
	AutoMigrate bool
}

type HTTP struct {
	IP   string
	Port string
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
		IP:   "0.0.0.0",
		Port: "8818",
	})

	viper.SetDefault("db", DB{
		Driver:      "postgresql",
		DSN:         "host=localhost user=postgres password=business dbname=business port=5432 sslmode=disable TimeZone=Asia/Shanghai",
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
		Use:       false,
		KeyPrefix: "system_manage",
	})

	viper.SetDefault("rpc", RPC{
		Host: "0.0.0.0",
		Port: 9060,
	})

	viper.SetDefault("nacos", Nacos{
		Register: false,
		ClientIP: "127.0.0.1",
		Host:     "127.0.0.1",
		Port:     10948,
		Weight:   0,
	})
	viper.SetDefault("enablegraph", false)
}

// set 设置一些固定不可改变的配置
func set() {

}

// doFlagConfig 绑定终端输入
func doFlagConfig() (configPath string) {
	// 配置文件路径
	pflag.StringVarP(&configPath, "config", "c", "config/config.yaml", "")
	pflag.String("http.ip", "localhost", "")
	pflag.Int("http.port", 8818, "")
	// 数据库
	pflag.String("db.driver", "postgresql", "")
	pflag.String("db.dsn", "host=localhost user=postgres password=business dbname=business port=5432 sslmode=disable TimeZone=Asia/Shanghai", "")
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	return
}

// doEnvConfig 	// 绑定环境变量
func doEnvConfig() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}

func InitConfigs() (cfg *Configs, err error) {
	var (
		confPath        = doFlagConfig()
		globalCfg       = Get()
		existConfigFile bool
	)

	doEnvConfig()
	setDefault()
	if _, _err := os.Stat(confPath); _err != nil {
		base, _ := path.Split(confPath)
		if _, _err := os.Stat(base); _err != nil {
			_ = os.MkdirAll(base, os.ModePerm)
		}

	} else {
		// 载入配置
		viper.SetConfigFile(confPath)
		viper.SetConfigType("yaml")
		existConfigFile = true
		//搜索配置文件，获取配置
		if err = viper.ReadInConfig(); err != nil {
			return
		}
	}
	if err = viper.Unmarshal(globalCfg); err != nil {
		return
	}
	if !existConfigFile {
		if err = viper.SafeWriteConfigAs(confPath); err != nil {
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
