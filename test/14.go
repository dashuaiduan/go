package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonTime time.Time

// 实现它的json序列化方法
func (this JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Student1 struct {
	Name  string   `json: "name" `
	Brith JsonTime `json: "brith" `
}

func main() {

	stu1 := Student1{
		Name:  "qiangmzsx",
		Brith: JsonTime(time.Date(1993, 1, 1, 20, 8, 23, 28, time.Local)),
	}
	b1, err := json.Marshal(stu1.Brith)
	if err != nil {
		println(err)
	}

	println(string(b1)) //{"name":"qiangmzsx","brith":"1993-01-01 20:08:23"}
}
