package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func main() {

	was := os.Args
	if len(was) == 1 {
		return
	}
	t2, err := strconv.Atoi(was[1])
	if err != nil {
		return
	}

	var waitGroup sync.WaitGroup
	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < t2; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Print(x)
		}(i)
	}
	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting ...")
}
