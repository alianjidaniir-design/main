package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/alianjidaniir-design/sqlite06"
)

var Min = 'A'
var Max = 'z'

func random(r *rand.Rand) byte {
	return byte(r.Intn(int(Max-Min+1)) + int(Min))
}

func getString(r *rand.Rand, length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(random(r))
	}
	return sb.String()
}

func main() {

	sqlite06.Filename = "ch01.db"
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
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	randomUsername := strings.ToLower(getString(r, 5))
	t := sqlite06.Userdata{
		Username:    randomUsername,
		Name:        "Ali",
		Surname:     "Anjidani",
		Description: "This is for me"}

	fmt.Println("Adding  username", randomUsername)
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
	randomUsername = strings.ToLower(getString(r, 5))
	randomName := getString(r, 7)
	randomSurname := getString(r, 10)
	fmt.Println(*r)
	dsc := time.Now().Format("2006-01-02 15:04:05")
	t = sqlite06.Userdata{
		Username:    randomName,
		Name:        randomName,
		Surname:     randomSurname,
		Description: dsc,
	}
	id = sqlite06.AddUser(t)
	if id == -1 {
		fmt.Println("There was an error adding user", t.Username)
	}

	dsc = time.Now().Format("2006-01-02 15:04:05")
	t.Description = dsc
	err = sqlite06.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
	}

}
