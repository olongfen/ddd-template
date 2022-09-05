package conf

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
	Addr string
	Host string
	Port int
}

func Get() *Configs {
	return Confs
}

var Confs = new(Configs)
