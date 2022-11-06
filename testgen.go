package main

import (
	"fmt"
)

//go:generate go run ./testgen.go
//go:generate go test -v ./testdash/...

func main() {
	fmt.Println("rbtyang test go:generate ok")
}
