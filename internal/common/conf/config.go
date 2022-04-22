package conf

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
	PEMFile string
	KeyFile string
	TLS     bool
	Host    string
	Port    int
	Clients []GRPCClient
}

type GRPCClient struct {
	AppID  string
	AppKey string
}

type Server struct {
	Http HTTP
	GRpc GRpc
}

func Get() *Configs {
	return Confs
}

var Confs = new(Configs)
