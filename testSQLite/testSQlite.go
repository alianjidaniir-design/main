package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//connect or create ac sqLite database
	db, err := sql.Open("sqlite3", "testt.db")
	if err != nil {
		fmt.Println("Error connection:", err)
		return
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		fmt.Println("Version : ", err)
		return
	}
	fmt.Println("SQlite3 version", version)

	err = os.Remove("test.db")
	if err != nil {
		return
	}
}
