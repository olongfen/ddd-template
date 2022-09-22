package cmd

import (
	"ddd-template/internal/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "ddd-template",
	Short: "application command",
}

func init() {
	cobra.OnInitialize(config.InitConfigs)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "./configs/config.yaml", "config file "+
		"(default is ./configs/config.yaml)")
	
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		zap.L().Fatal(err.Error())
		os.Exit(1)
	}
}
