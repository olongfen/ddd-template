package config

import (
	"bytes"
	"e.coding.net/zkxrsz/starwiz/zkxr_center_backend/system-manage/pkg/utils"
	"github.com/fsnotify/fsnotify"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Configs struct {
	WatchConfig bool
	HTTP        HTTP
	RPC         RPC
	Database    Database
	Languages   []string
	Log         Log
	Redis       Redis
	JWT         JWT
	Nacos       Nacos
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

type JWT struct {
	Auth                   bool
	AccessTokenPrivateKey  string
	AccessTokenPublicKey   string
	RefreshTokenPrivateKey string
	RefreshTokenPublicKey  string
	AccessTokenExpiresIn   int // 单位分钟
	RefreshTokenExpiresIn  int // 单位分钟
	AccessTokenMaxAge      int
	RefreshTokenMaxAge     int
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
	Port int
}

type Redis struct {
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
		Port: 8818,
	})
	viper.SetDefault("languages", []string{"cn", "en"})
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
		KeyPrefix: "system_manage",
	})

	viper.SetDefault("rpc", RPC{
		Host: "0.0.0.0",
		Port: 9060,
	})

	viper.SetDefault("jwt", JWT{
		AccessTokenPrivateKey: `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIM/VQnvP/2h9De/P5GyLcpk8VjcQKwBGSn873vu5orOyoAoGCCqGSM49
AwEHoUQDQgAE3hXUV+yvG0aq+NMOiU/LqSdSwBuMyIkOBonfL6885mW+1nE7cC2J
HfwkUPILMgLePSnSldMFTij2fb6m2ABjQA==
-----END EC PRIVATE KEY-----`,
		AccessTokenPublicKey: `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE3hXUV+yvG0aq+NMOiU/LqSdSwBuM
yIkOBonfL6885mW+1nE7cC2JHfwkUPILMgLePSnSldMFTij2fb6m2ABjQA==
-----END PUBLIC KEY-----`,
		RefreshTokenPrivateKey: `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIBn0pbZ6veqGJRC88IR1cRWbIYXLK/bQV3D3WBgipexRoAoGCCqGSM49
AwEHoUQDQgAEUgryp8sKrfu1et3nxObupVLoS5l+RedxwsjkH3EvaTY120g0MA1e
kwAPdEQ8NVBZ/e2Rulm7pPAesPhstdonZg==
-----END EC PRIVATE KEY-----
`,
		RefreshTokenPublicKey: `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUgryp8sKrfu1et3nxObupVLoS5l+
RedxwsjkH3EvaTY120g0MA1ekwAPdEQ8NVBZ/e2Rulm7pPAesPhstdonZg==
-----END PUBLIC KEY-----
`,
		AccessTokenExpiresIn:  15,
		RefreshTokenExpiresIn: 60,
		AccessTokenMaxAge:     15,
		RefreshTokenMaxAge:    60,
	})
	viper.SetDefault("nacos", Nacos{
		Register:   false,
		ClientName: "127.0.0.1",
		Host:       "127.0.0.1",
		Port:       10948,
		Weight:     0,
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
			originalBytes []byte
			changeBytes   []byte
		)
		// 读取旧文件含有的配置
		if originalBytes, err = os.ReadFile(confPath); err != nil {
			log.Fatalln("ReadFile", err)
		}
		if err = utils.Copier(viper.AllSettings(), globalCfg); err != nil {
			log.Fatalln("Copier", err)
		}
		if err = yaml.Unmarshal(originalBytes, globalCfg); err != nil {
			log.Fatalln("Unmarshal", err)
		}
		if changeBytes, err = jsoniter.Marshal(globalCfg); err != nil {
			log.Fatalln("Marshal", err)
		}
		if err = viper.ReadConfig(bytes.NewBuffer(changeBytes)); err != nil {
			log.Fatalln("ReadConfig from changeBytes", err)
		}
		if err = viper.WriteConfig(); err != nil {
			log.Fatalln("WriteConfig", err)
		}

	}
	if globalCfg.WatchConfig {
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
			if err = viper.Unmarshal(globalCfg); err != nil {
				log.Fatal(err)
			}
		})
	}

	log.Println("config init success")
	return globalCfg
}
