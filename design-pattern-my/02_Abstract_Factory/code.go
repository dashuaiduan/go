package _2_Abstract_Factory

//抽象工厂模式
import (
	"errors"
	"fmt"
)

// 车
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

// 飞机
type Aircraft interface {
	didi()
}

type public_aircraft struct {
}

func (p public_aircraft) didi() {
	fmt.Println("public_aircraft is didi....")
}

type private_aircraft struct {
}

func (p private_aircraft) didi() {
	fmt.Println("private_aircraft is didi....")
}

// 抽象为交通工具
type VehicleFactory interface {
	GetCar(s string) (error, Car)
	GetAircraft(s string) (error, Aircraft)
}

// 飞机工厂
type AircraftFactory struct {
}

func (a AircraftFactory) GetCar(s string) (error, Car) {
	return nil, nil
}
func (a AircraftFactory) GetAircraft(s string) (error, Aircraft) {
	if s == "private_aircraft" {
		return nil, new(private_aircraft)
	}
	if s == "public_aircraft" {
		return nil, new(public_aircraft)
	}
	return errors.New("fail"), nil
}

type CarFactory struct {
}

func (c CarFactory) GetCar(s string) (error, Car) {
	if s == "bmw" {
		return nil, new(bmw)
	}
	if s == "porsche" {
		return nil, new(porsche)
	}
	return errors.New("fail"), nil
}
func (c CarFactory) GetAircraft(s string) (error, Aircraft) {
	return nil, nil
}

func NewFactory(s string) (error, VehicleFactory) {
	if s == "car" {
		return nil, new(CarFactory)
	}
	if s == "air" {
		return nil, new(AircraftFactory)
	}
	return nil, nil
}
