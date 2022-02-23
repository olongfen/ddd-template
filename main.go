package main

import (
	_ "ddd-template/docs"
	"ddd-template/infra/conf"
	"ddd-template/infra/xlog"
	"flag"
	"go.uber.org/zap"
)

var (
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	conf.InitConf(flagconf)
	cfg := conf.Get()
	var (
		logger *zap.Logger
	)
	if cfg.Environment == "dev" {
		logger = xlog.NewDevelopment()
	} else {
		logger = xlog.NewProduceLogger()
	}
	xhttp := Init(cfg, logger)
	xhttp.Router().Start()
}
