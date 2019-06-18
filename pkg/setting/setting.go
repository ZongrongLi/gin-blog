package setting

import (
	"time"

	"github.com/golang/glog"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string
}

var ServerSetting = &Server{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var AppSetting = &App{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Log struct {
	Level  string
	Path   string
	Stderr string
}

var LogSetting = &Log{}

func Setup() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		glog.Errorf("Fail to parse 'conf/app.ini': %v", err)

	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		glog.Errorf("Cfg.MapTo AppSetting err: %v", err)
	}

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		glog.Errorf("Cfg.MapTo ServerSetting err: %v", err)

	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		glog.Errorf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("log").MapTo(LogSetting)
	if err != nil {
		glog.Errorf("Cfg.MapTo ServerSetting err: %v", err)
	}
}

// var (
// 	Cfg *ini.File

// 	RunMode string

// 	LOG_LEVEL  string
// 	LOG_PATH   string
// 	LOG_STDERR string

// 	HTTPPort     int
// 	ReadTimeout  time.Duration
// 	WriteTimeout time.Duration

// 	PageSize  int
// 	JwtSecret string
// )

// func init() {
// 	var err error
// 	Cfg, err = ini.Load("conf/app.ini")
// 	if err != nil {
// 		glog.Errorf("Fail to parse 'conf/app.ini': %v", err)
// 	}

// 	LoadLog()
// 	LoadBase()
// 	LoadServer()
// 	LoadApp()
// }
// func LoadLog() {
// 	sec, err := Cfg.GetSection("glog")
// 	if err != nil {
// 		glog.Errorf("Fail to get section 'glog': %v", err)
// 	}
// 	LOG_LEVEL = sec.Key("LEVEL").MustString("0")
// 	LOG_PATH = sec.Key("PATH").MustString("./glog")
// 	LOG_STDERR = sec.Key("STDERR").MustString("true")
// }

// func LoadBase() {
// 	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
// }

// func LoadServer() {
// 	sec, err := Cfg.GetSection("server")
// 	if err != nil {
// 		glog.Errorf("Fail to get section 'server': %v", err)
// 	}

// 	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

// 	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
// 	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
// 	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
// }

// func LoadApp() {
// 	sec, err := Cfg.GetSection("app")
// 	if err != nil {
// 		glog.Errorf("Fail to get section 'app': %v", err)
// 	}

// 	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
// 	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
// }
