package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//connect or create ac sqLite database
	db, err := sql.Open("sqlite3", "ch0code1.db")
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
	fmt.Println("SQLite3 version", version)

	os.Remove("ch0code1.db")

}
