package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"fmt"

	"github.com/tiptok/gonat/core"
	"github.com/tiptok/gonat/global"
)

var host core.Host
var exit chan int

func main() {
	defer func() {
		exit <- 1 //异常退出
	}()
	runtime.GOMAXPROCS(runtime.NumCPU())
	host = core.Host{}
	host.Start(global.Param.Protocol)
	//等待退出

	//pprof
	go func() {
		http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", global.Param.PProfPort), nil)
	}()
	<-exit
}
