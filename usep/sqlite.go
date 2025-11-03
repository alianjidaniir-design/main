package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/alianjidaniir-design/sqilite06"
)

var Min = -12
var Max = 90

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getstring(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(Min, Max)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {

	sqilite06.Filename = "ch06.db"

}
