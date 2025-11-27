package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	aew := os.Args
	if len(aew) == 1 {
		fmt.Println("Usage: ./main aew")
		return
	}
	temp, err := strconv.Atoi(aew[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Creating %d gor.\n", temp)
	for i := 2; i <= temp; i++ {
		go func(x int) {
			fmt.Printf("%d", x)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("\nExiting...")
}
