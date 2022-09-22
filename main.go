package main

import (
	"context"
	"ddd-template/internal/config"
	"ddd-template/internal/ports"
	"ddd-template/pkg/xlog"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func main() {
	var (
		server    ports.HttpServer
		ctx       = context.Background()
		logger, _ = zap.NewProduction()
		fc        func()
	)
	execute()
	cfg := config.InitConfigs(configFile)
	if cfg.Log.Debug {
		logger = xlog.NewDevelopment()
	} else {
		logger = xlog.NewProduceLogger()
	}
	xlog.Log = logger
	server, fc = NewServer(ctx, cfg, logger)
	defer fc()
	ports.RunHTTPServer(func(app2 *fiber.App) *fiber.App {
		return ports.HandlerFromMux(server, app2)
	}, cfg.HTTP, logger)

}

var configFile string

var rootCmd = &cobra.Command{
	Use:   "ddd-template",
	Short: "application command",
	Long:  "application command,exec some action",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "./configs/config.yaml", "config file "+
		"(default is ./configs/config.yaml)")

}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		zap.L().Fatal(err.Error())
		os.Exit(1)
	}
}
