# Notes - Hello World

📅 Date: 28.10.2025
🔢 Go Version: 1.23
🏗️ Folder: 00-basics/hello-world

---

## 🎯Goal
Set up Go 1.23 environment and run the first program to understand basic project structure, package declaration, the `main` function and imports.

---

## 🧠 Key Learnings
- Every Go executable starts from `package main` and `func main()`.
- Imports must be explicit — unused imports cause compile errors.
- `go run .` compiles and runs code directly; `go build` creates a binary.
- Go enforces code formatting with `gofmt`; no stylistic debates.
- The standard library (`fmt`, `runtime`, `time`) already covers most basic needs.

---

## 📋 Code Summary
```go
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

```
