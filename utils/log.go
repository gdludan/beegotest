package utils

// 导入beego日志模块 写日志到logs目录
import (
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// 获取当前时间
func GetNowTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func init() {
	log := logs.NewLogger()
	log.Async()
	// 写入到logs目录下
	log.SetLogger("file", `{"filename":"logs/"+GetNowTime()+".log"}`)
}
