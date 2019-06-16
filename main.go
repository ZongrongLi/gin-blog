package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/golang/glog"

	"github.com/tiancai110a/gin-blog/pkg/setting"
	"github.com/tiancai110a/gin-blog/routers"
)

func init() {
	flag.Set("alsologtostderr", setting.LOG_STDERR) // 日志写入文件的同时，输出到stderr
	flag.Set("log_dir", setting.LOG_PATH)           // 日志文件保存目录
	flag.Set("v", setting.LOG_LEVEL)                // 配置V输出的等级。
	flag.Parse()

}

func main() {
	router := routers.InitRouter()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	go func(c chan os.Signal) {
		<-c
		glog.Flush()
		os.Exit(0)
	}(c)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
