package main

import (
	"fmt"
	"github.com/PichuChen/daemon"
	_ "github.com/PichuChen/ifconfig/routers"
	"github.com/astaxie/beego"
	"os"
	"os/signal"
	"syscall"
)

const (
	HostVar = "VCAP_APP_HOST"
	PortVar = "VCAP_APP_PORT"

	name        = "ifconfig.tw"
	description = "ifconfig.tw daemon"
)

type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (string, error) {
	usage := "Usage: " + os.Args[0] + " install | remove | start | stop | status"

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	fmt.Printf("HTTP Port: %d\n", beego.BConfig.Listen.HTTPPort)
	go func() {
		// beego.Run(beego.BConfig.Listen.HTTPAddr + ":" + fmt.Sprintf("%d", beego.BConfig.Listen.HTTPPort)) // Block
		beego.Run()
		fmt.Println("beego exit")
		interrupt <- os.Kill
	}()
	<-interrupt
	// good bye
	return "Server stop", nil

}

func main() {
	srv, err := daemon.New(name, description, []string{}...)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		fmt.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)

}
