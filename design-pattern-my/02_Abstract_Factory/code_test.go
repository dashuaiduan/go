package _2_Abstract_Factory

import (
	"testing"
)

func Test1(t *testing.T) {
	//var car Car

	ok, airFactory := NewFactory("air")
	if ok == nil {
		var air Aircraft
		ok1, air := airFactory.GetAircraft("private_aircraft")
		if ok1 == nil {
			air.didi()
		}
		ok2, air := airFactory.GetAircraft("public_aircraft")
		if ok2 == nil {
			air.didi()
		}
	}

	ok, carFactory := NewFactory("car")
	if ok == nil {
		var car Car
		ok1, car := carFactory.GetCar("bmw")
		if ok1 == nil {
			car.run()
		}
		ok2, car := carFactory.GetCar("porsche")
		if ok2 == nil {
			car.run()
		}
	}
}
