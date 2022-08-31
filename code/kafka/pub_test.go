package kafka

import (
	"fmt"
	"strings"
	"testing"
)

var MqQueue Client

func init() {
	//	初始化客户端
	MqQueue = Client{
		Name:     "kafka",
		Server:   strings.Split("ckafka-lw3ar7oe.ap-guangzhou.ckafka.tencentcloudmq.com:6012", "::"),
		ClientID: "clientIdfffffffffff",
		User:     "ckafka-lw3ar7oe#worthcloud",
		Password: "Aa123456",
	}
	if err := MqQueue.NewClient(); err != nil {
		panic("MqQueue.NewClient() failed " + err.Error())
	}
}

func TestName(t *testing.T) {
	//init1()
	msg := MsgPub{
		Topic: "CN_DEV_IOT_DEVICE_STATE", // 自定义，必须提前创建
		Data:  []byte("我来啦！！！"),
	}
	err := MqQueue.Publish(&msg)
	fmt.Println(err)

}
