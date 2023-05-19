package main

import (
	"ddd-template/internal/service"
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
	// 创建服务
	if server, cleanup, err = NewServer(); err != nil {
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
