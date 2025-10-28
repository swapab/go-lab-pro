package main

import "fmt"

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

func Start(d Drivable) {
	d.Drive()
}

func main() {
	myCar := Vehicle{Brand: "Audi", Model: "Q4", Year: 2024}
	fmt.Println(myCar.Greet())
	Start(myCar)
}
