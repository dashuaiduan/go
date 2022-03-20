package _1_Factory

//工厂模式
import (
	"errors"
	"fmt"
)

// 如果是抽象工厂类 ，则将car再提升一下，比如多一个接口交通工具接口 ，有一个飞机接口 ，如下 代码再来一套飞机的代码
type Car interface {
	run()
}

type bmw struct {
}

func (b bmw) run() {
	fmt.Println("bmw is running....")
}

type porsche struct {
}

func (p porsche) run() {
	fmt.Println("bmw is running....")
}

func NewCar(s string) (error, Car) {
	if s == "bmw" {
		return nil, new(bmw)
	}
	if s == "porsche" {
		return nil, new(porsche)
	}
	return errors.New("fail"), nil
}
