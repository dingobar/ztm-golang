//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Truck
)

func (vt VehicleType) String() string {
	switch vt {
	case Motorcycle:
		return "motorcycle"
	case Car:
		return "car"
	case Truck:
		return "truck"
	default:
		return "unknown vehicle type"
	}
}

type Liftable interface {
	Lift()
}

type Vehicle struct {
	Type  VehicleType
	Model string
}

func (v *Vehicle) Lift() {
	fmt.Println("A", v.Type, "of model", v.Model, "was lifted!")
}

func main() {
	car := Vehicle{Car, "Ford Mondeo"}
	motorcycle := Vehicle{Motorcycle, "Kawazaki"}
	truck := Vehicle{Truck, "Ford F-150"}
	car.Lift()
	motorcycle.Lift()
	truck.Lift()
}
