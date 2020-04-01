package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ajdwfnhaps/easy-go/plugin/web"
)

func main() {

	//使用 easy-gin
	web.UseEasyGinPlugin("conf/config.toml")
	handleSignal()
}

func handleSignal() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case <-c:
		fmt.Println("服务退出")
	}
}
