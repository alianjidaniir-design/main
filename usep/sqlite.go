package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/alianjidaniir-design/sqlite06"
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

	sqlite06.Filename = "ch06.db"
	data, err := sqlite06.ListUsers()
	if err != nil {
		fmt.Println("ListUsers():", err)
		return
	}
	if len(data) != 0 {
		for _, v := range data {
			fmt.Println(v)
		}
	}
	random_username := strings.ToLower(getstring(5))
	t := sqlite06.Userdata{
		Username:    random_username,
		Name:        "Ali",
		Surname:     "Anjidani",
		Description: "This is for me",
	}
	fmt.Println("Adding  username", random_username)
	id := sqlite06.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}
	err = sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println("DeleteUser:", err)
	} else {
		fmt.Println("User with ID:", id, "Deleted!")
	}
	err = sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println("DeleteUser:", err)
	}
	random_username = strings.ToLower(getstring(5))
	random_name := getstring(7)
	random_surname := getstring(10)
	dsc := time.Now().Format("2006-01-02 15:04:05")
	t = sqlite06.Userdata{
		Username:    random_name,
		Name:        random_name,
		Surname:     random_surname,
		Description: dsc,
	}
	id = sqlite06.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
		fmt.Println(-1)
	}

}
