package main

import (
	"ddd-template/common/conf"
	"ddd-template/common/xlog"
	_ "ddd-template/docs"
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
	app := Init(cfg, logger)
	app.Handles("/api/v1")
	app.Run(conf.Get().Server.Http.Addr)
}
