package main

import (
	_ "embed"
	"fmt"
)

//go:embed printsource.go
var src string

func main() {
	fmt.Print(src)
}
