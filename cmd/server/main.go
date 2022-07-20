package main

import (
	"ddd-template/internal/app"
	"ddd-template/internal/common/xlog"
	"github.com/urfave/cli/v2"
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
	var application *app.Application
	defer func() {
		_ = xlog.Log.Sync()
		_ = application.Close()
	}()
	app, _ := NewServer(flagconf)
	app.Run()
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
