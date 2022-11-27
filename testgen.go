package main

import (
	"fmt"
)

//go:generate go test ./...
//go:generate go test -v ./testdash/...
//go:generate go run ./testgen.go

func main() {
	fmt.Println("rbtyang test go:generate ok")
}
