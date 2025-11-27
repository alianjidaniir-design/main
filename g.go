package main

import (
	"fmt"
	"runtime"
	"time"
)

func pm(xx int) {
	fmt.Println(xx*3*xx, xx)
	return
}

func main() {
	fmt.Println("runtime.Compiler,", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOOS, "machine")
	fmt.Println(runtime.Version())
	fmt.Printf("GOMAXPROCS=%d\n", runtime.GOMAXPROCS(1))

	go func(xx int) {
		fmt.Println(xx)
	}(5)
	go pm(13)
	go pm(14)
	time.Sleep(time.Second * 1)
	fmt.Println("Exiting")

}
