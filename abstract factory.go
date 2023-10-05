package main

import "fmt"

type VehicleFactory interface {
	CreateCar() Car
	CreateMotorcycle() Motorcycle
}

type Car interface {
	Drive() string
}

type SportsCar struct{}

func (c SportsCar) Drive() string {
	return "Driving a sports car"
}

type Motorcycle interface {
	Ride() string
}

type Cruiser struct{}

func (m Cruiser) Ride() string {
	return "Riding a cruiser motorcycle"
}

type FordFactory struct{}

func (f FordFactory) CreateCar() Car {
	return SportsCar{}
}

func (f FordFactory) CreateMotorcycle() Motorcycle {
	return Cruiser{}
}

type HondaFactory struct{}

func (f HondaFactory) CreateCar() Car {
	return nil
}

func (f HondaFactory) CreateMotorcycle() Motorcycle {
	return Cruiser{}
}

func main() {
	fordFactory := FordFactory{}

	fordCar := fordFactory.CreateCar()
	fordMotorcycle := fordFactory.CreateMotorcycle()

	if fordCar != nil {
		fmt.Println(fordCar.Drive())
	} else {
		fmt.Println("Ford does not make cars in this example")
	}
	fmt.Println(fordMotorcycle.Ride())

	hondaFactory := HondaFactory{}

	hondaCar := hondaFactory.CreateCar()
	hondaMotorcycle := hondaFactory.CreateMotorcycle()

	if hondaCar != nil {
		fmt.Println(hondaCar.Drive())
	} else {
		fmt.Println("Honda does not make cars in this example")
	}
	fmt.Println(hondaMotorcycle.Ride())
}
