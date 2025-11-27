package main

import (
	"fmt"
	"runtime"
	"time"
)

func pm(xx int) int {
	return xx * xx
}

func main() {
	fmt.Println("runtime.Compiler,", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOOS, "machine")
	fmt.Println(runtime.Version())
	fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(1))

	go func(x int) {
		fmt.Println(123)
	}(5)
	go pm(13)
	go pm(14)
	time.Sleep(time.Second)
	fmt.Println("Exiting")

}
