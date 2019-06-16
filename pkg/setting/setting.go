package setting

import (
	"time"

	"github.com/golang/glog"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	LOG_LEVEL  string
	LOG_PATH   string
	LOG_STDERR string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		glog.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadLog()
	LoadBase()
	LoadServer()
	LoadApp()
}
func LoadLog() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		glog.Fatalf("Fail to get section 'log': %v", err)
	}
	LOG_LEVEL = sec.Key("LEVEL").MustString("0")
	LOG_PATH = sec.Key("PATH").MustString("./log")
	LOG_STDERR = sec.Key("STDERR").MustString("true")
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		glog.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		glog.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
