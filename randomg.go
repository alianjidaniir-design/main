package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func random(max, min int) int {
	return rand.Intn(max-min) + min
}
func createfile(file string) {
	_, err := os.Stat(file)
	if err != nil {
		fmt.Println("File does not exist")
		return
	}

	f, err := os.Create(file)
	if err != nil {
		fmt.Println("Error creating file")
	}

	lines := random(1, 54)

	for i := 0; i < lines; i++ {
		data := random(3, 47)
		fmt.Fprintln(f, "%d\n", data)
	}

}

func main() {

	as := os.Args
	if len(as) == 1 {
		fmt.Println("Usage: randomFiles firstInt lastInt filename directory")
		return
	}

	start, err := strconv.Atoi(as[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	end, err := strconv.Atoi(as[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	if start > end {
		fmt.Println(end, "<", start)
		return
	}

	filename := as[3]
	path := as[4]
	_, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup
	for i := start; i < end; i++ {
		waitGroup.Add(1)
		go func(i int) {
			filepath := filepath.Join(path, fmt.Sprintf("%s%D", filename, i))
			defer waitGroup.Done()
			createfile(filepath)
		}(i)

	}

	waitGroup.Wait()

}
