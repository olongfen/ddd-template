package main

import (
	"ddd-template/internal/service"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// @title           documents
// @version         1.0
// @description     用户管理系统api文档
// @contact.name   olongfen
// @contact.email  olongfen@gmail.com
// @schemes http https
// @BasePath  /
// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	var (
		server  *service.Server
		cleanup func()
		err     error
		wg      = sync.WaitGroup{}
		done    = make(chan struct{})
	)
	// 监听关闭
	setupCloseHandler(done)
	execute()
	// 创建服务
	if server, cleanup, err = NewServer(configFile); err != nil {
		log.Fatalln("NewServer", err)
	}
	go func() {
		for range done {
			cleanup()
			log.Println("end of process ")
			os.Exit(0)
		}
	}()
	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
			if r := recover(); r != nil {
				done <- struct{}{}
			}
		}()
		server.Http.Run()
	}()

	wg.Wait()
}

func setupCloseHandler(done chan struct{}) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		done <- struct{}{}
		log.Println("Ctrl+C pressed in Terminal")
	}()
}

var configFile string

var rootCmd = &cobra.Command{
	Use:   "system-manage",
	Short: "system-manage command",
	Long:  "system-manage command,exec some action",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "./configs/config.yaml", "config file "+
		"(default is ./configs/config.yaml)")

}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
