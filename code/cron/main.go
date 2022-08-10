package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 这个包在 test中 无效， 官方文档中的6个*  秒级任务 暂时没有测试通过  分钟5个* 是有效的
	c := cron.New()
	c.AddFunc("*/1 * * * *", task)
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()
	c.AddFunc("*/1 * * * *", func() { fmt.Println("start 之后启动的") })
	//c.Stop()  // Stop the scheduler (does not stop any jobs already running).
	time.Sleep(time.Minute * 30)
}
func task() {
	fmt.Println("adff")
}
