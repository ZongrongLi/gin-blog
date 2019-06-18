package log

import (
	"flag"

	"github.com/tiancai110a/gin-blog/pkg/setting"
)

func Setup() {
	flag.Set("alsologtostderr", setting.LogSetting.Stderr)
	flag.Set("log_dir", setting.LogSetting.Path)
	flag.Set("v", setting.LogSetting.Level)
	flag.Parse()
}
