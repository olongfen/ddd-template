package main

import (
	"ddd-template/common/conf"
	"ddd-template/common/xlog"
	cli "github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"log"
	"os"
)

const (
	confFlag = "conf"
)

var (
	initFlags = []cli.Flag{
		&cli.StringFlag{
			Name:     confFlag,
			Aliases:  []string{"c"},
			Usage:    "config path, eg: -c config.yaml",
			Required: false,
			Value:    "./configs/config.yaml",
		},
	}
)

func initAction(c *cli.Context) (err error) {
	var flagconf = c.String(confFlag)
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
	xlog.Log = logger
	app := Init(cfg, logger)
	app.Run(conf.Get().Server)
	return
}

func main() {
	app := cli.NewApp()
	app.Action = initAction
	app.Flags = initFlags
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
