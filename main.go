package main

import (
	"flag"
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"

	"github.com/golang/glog"

	"github.com/tiancai110a/gin-blog/pkg/setting"
	"github.com/tiancai110a/gin-blog/routers"
)

func init() {
	flag.Set("alsologtostderr", setting.LOG_STDERR) // 日志写入文件的同时，输出到stderr
	flag.Set("log_dir", setting.LOG_PATH)           // 日志文件保存目录
	flag.Set("v", setting.LOG_LEVEL)                // 配置V输出的等级。
	glog.Info("================================================", setting.LOG_LEVEL)
	flag.Parse()

}

func main() {
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, os.Interrupt, os.Kill)

	// go func(c chan os.Signal) {
	// 	<-c
	// 	glog.Flush()
	// 	os.Exit(0)
	// }(c)

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
