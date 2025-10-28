package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Go Version:", runtime.Version())
	fmt.Println("OS/Arch:", runtime.GOOS, runtime.GOARCH)
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("Current Time:", time.Now())
}
