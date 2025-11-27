package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.Compiler,", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOOS, "machine")
	fmt.Println(runtime.Version())
	fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(0))
}
