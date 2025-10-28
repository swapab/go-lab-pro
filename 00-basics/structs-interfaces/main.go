package main

import (
	"fmt"
)

type Vehicle struct {
	Brand string
	Model string
	Year  int
}

func (v Vehicle) Greet() string {
	return fmt.Sprintf("Welcome to your %s, %s from year %d", v.Brand, v.Model, v.Year)
}

type Drivable interface{ Drive() }

func (v Vehicle) Drive() {
	fmt.Println(v.Brand, "is now driving!")
}

type Bike struct {
	Brand string
}

func (v Bike) Drive() {
	fmt.Println(v.Brand, "bike is pedaling!")
}

type Truck struct {
	Brand string
}

func (v Truck) Drive() {
	fmt.Println(v.Brand, "truck is hauling!")
}

func Start(d Drivable) {
	d.Drive()
}

func main() {
	myCar := Vehicle{Brand: "Audi", Model: "Q4", Year: 2024}
	fmt.Println(myCar.Greet())
	Start(myCar)
	Start(Bike{Brand: "Hero"})
	Start(Truck{Brand: "Volvo"})
}
