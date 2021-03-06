package logurs1

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestLog(t *testing.T) {
	u := User{
		Name: "dj",
		Age:  18,
	}
	//内置log包 提供了三组函数：
	//log库只提供了三组接口，功能过于简单了。
	//Print/Printf/Println：正常输出日志；
	//Panic/Panicf/Panicln：输出日志后，以拼装好的字符串为参数调用panic；
	//Fatal/Fatalf/Fatalln：输出日志后，调用os.Exit(1)退出程序。
	log.SetPrefix("大帅日志: ")                                                      // 为每条日志文本前增加一个前缀。
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds | log.Llongfile) // 自定义log输出格式 原有的选项会被覆盖掉！

	log.Printf("%s login, age:%d", u.Name, u.Age)
	log.Panicf("Oh, system error when %s login", u.Name)
	log.Fatalf("Danger! hacker %s login", u.Name)
}

func TestLogrus(t *testing.T) {

	//地鼠文档  https://www.topgoer.cn/docs/goday/goday-1crg2adjknouc

	//Panic：记录日志，然后panic。
	//Fatal：致命错误，出现错误时程序无法正常运转。输出日志后，程序退出；
	//Error：错误日志，需要查看原因；
	//Warn：警告信息，提醒程序员注意；
	//Info：关键操作，核心流程的日志；
	//Debug：一般程序中输出的调试信息；
	//Trace：很细粒度的信息，一般用不到；

	logrus.SetLevel(logrus.TraceLevel) // 设置输出级别

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	//logrus.Fatal("fatal msg")
	//logrus.Panic("panic msg")

	logrus.SetReportCaller(true) //设置在输出日志中添加文件名和方法信息：

	//time="2022-04-22T09:40:09+08:00" level=info msg="info msg" func=test/code/logurs.TestLogrus file="D:/goproject/go/code/logurs/code_test.go:56"
	logrus.Info("info msg")

	fmt.Println("-------------------------------")
	// 固定输出信息  使用requestLogger 对象输出的时候 固定输出WithFields配置的内容，同时使用requestLogger会继承logrus的配置
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "192.168.32.15",
	})

	//level=info msg="info msg" func=test/code/logurs.TestLogrus file="D:/goproject/go/code/logurs/code_test.go:67" ip=192.168.32.15 user_id=10010
	requestLogger.Info("info msg")
	requestLogger.Error("error msg")
}

// 测试多重写MultiWriter
func TestWww(t *testing.T) {
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info msg")
}

//自定义log
func TestNew(t *testing.T) {
	log := logrus.New()

	log.SetLevel(logrus.InfoLevel)
	log.SetFormatter(&logrus.JSONFormatter{}) // 支持两种日志格式，文本和 JSON，默认为文本格式。可以通过logrus.SetFormatter设置日志格式

	log.Info("info msg")

}
