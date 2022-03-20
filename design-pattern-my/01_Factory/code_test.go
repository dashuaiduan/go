package _1_Factory

import (
	"testing"
)

func Test1(t *testing.T) {
	//fmt.Println("bmw" == "bmw")

	var car Car
	ok, car := NewCar("bmw")
	if ok == nil {
		car.run()
	}
	ok1, car1 := NewCar("porsche")
	if ok1 == nil {
		car1.run()
	}
}
