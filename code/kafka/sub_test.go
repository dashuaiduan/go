package kafka

import (
	"fmt"
	"testing"
	"time"
)

func TestName1(t *testing.T) {
	//设备上下线消息订阅
	go MqQueue.Subscribe(&MsgSub{
		Topic:      "CN_DEV_IOT_DEVICE_STATE",
		GroupId:    "device-report-dev1111test",
		ClientName: "ClientName",
		MsgHandle:  listenerThingStateMsg,
	})
	time.Sleep(time.Hour)
}

// test debug 运行 Sub然后断点当前 处理函数  run方式运行pub 即可进入断点
func listenerThingStateMsg(msgBytes []byte) {
	fmt.Println(msgBytes)
}
