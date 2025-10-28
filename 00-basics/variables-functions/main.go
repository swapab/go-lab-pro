package main

import "fmt"

const pi = 3.14159

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	name := "Audi"
	year := 2024
	fmt.Printf("%s Model Year: %d\n", name, year)
	fmt.Println("Addition:", add(5, 7))

	result, err := divide(10, 2)
	fmt.Println("Division Result:", result, "Error:", err)
}
