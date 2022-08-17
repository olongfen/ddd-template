package main

import (
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
	defer func() {
		_ = xlog.Log.Sync()
	}()
	NewServer(flagconf)
	return
}

// @title demo
// @version 1.0
// @description demo
// @Schemes HTTP HTTPS
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @Accept application/json
// @Produce application/json
// @contact.name olongfen
// @contact.email olongfen@gmail.com
// @BasePath /api/v1
func main() {
	app := cli.NewApp()
	app.Action = initAction
	app.Flags = initFlags
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
