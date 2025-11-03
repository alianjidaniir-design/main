package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var dbname = "ch01.db"

func insertData(db *sql.DB, dsc string) error {

	cT := time.Now().Format(time.RFC822Z)
	stmt, err := db.Prepare("INSERT INTO book VALUES(NULL,?,?);")
	if err != nil {
		fmt.Println("Insert data table:", err)
		return err
	}
	_, err = stmt.Exec(cT, dsc)
	if err != nil {
		fmt.Println("Insert data table :", err)
		return err
	}
	return nil
}

func selectData(db *sql.DB, n int) error {
	rows, err := db.Query("SELECT * from book WHERE id > ? ", n)
	if err != nil {
		fmt.Println("Select1:", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var dt string
		var description string
		err := rows.Scan(&id, &dt, &description)
		if err != nil {
			fmt.Println("Row:", err)
			return err
		}
		date, err := time.Parse(time.RFC822Z, dt)
		if err != nil {
			fmt.Println("Date1:", err)
			return err
		}
		fmt.Printf("%d %s %s\n", id, date, description)
	}
	return nil
}

func main() {

	os.Remove(dbname)

	db, err := sql.Open("sqlite3", dbname)
	if err != nil {
		fmt.Println("Error connection:", err)
		return
	}

	defer db.Close()

	const create string = `
CREATE TABLE IF NOT EXISTS book (  
 id INTEGER NOT NULL PRIMARY KEY,
time TEXT NOT NULL,
 description TEXT);`

	_, err = db.Exec(create)
	if err != nil {
		fmt.Println("create table", err)
		return
	}
	for i := 1; i <= 10; i++ {
		dsc := "Description: " + strconv.Itoa(i)
		err = insertData(db, dsc)
		if err != nil {
			fmt.Println("Insert data:", err)
		}
	}
	// Select multiple rows
	err = selectData(db, 5)
	if err != nil {
		fmt.Println("Select2:", err)
	}
	time.Sleep(time.Second)
	//update data

	cT := time.Now().Format(time.RFC822Z)
	db.Exec("UPDATE book SET time = ? WHERE id > ?", cT, 7)

	err = selectData(db, 8)
	if err != nil {
		fmt.Println("Select3:", err)
		return
	}
	//Delete data

	stmt, err := db.Prepare("DELETE from book WHERE id = ?")
	_, err = stmt.Exec(8)
	if err != nil {
		fmt.Println("Delete:", err)
		return
	}

	err = selectData(db, 7)
	if err != nil {
		fmt.Println("Select4:", err)
		return
	}

	query, err := db.Query("SELECT count(*) as count from book WHERE id > ?", 7)
	if err != nil {
		fmt.Println("Select5:", err)
		return
	}
	defer query.Close()
	count := -100
	for query.Next() {
		_ = query.Scan(&count)
	}
	fmt.Println("count(*)", count)
}
