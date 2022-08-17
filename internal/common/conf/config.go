package conf

type Configs struct {
	HTTP        HTTP
	Database    Database
	Environment string
	Language    string
	Log         Log
}

type Log struct {
	Filename   string
	ErrorFile  string
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

func Get() *Configs {
	return Confs
}

var Confs = new(Configs)
