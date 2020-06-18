package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Profile struct
type Profile struct {
	ProfileID int
	FirstName string
	LastName  string
	Age       int
}

func main() {
	db, err := sql.Open("sqlite3", "./names.db")
	checkError(err)

	var profile Profile
	rows, err := db.Query("select ProfileID, FirstName, LastName, Age from profile")
	// rows, err := db.Query("select ProfileID, FirstName, LastName, Age from profile where ProfileID = ?", 2)
	checkError(err)

	for rows.Next() {
		err := rows.Scan(&profile.ProfileID, &profile.FirstName, &profile.LastName, &profile.Age)
		checkError(err)
		fmt.Println(profile)
	}

	rows.Close()
	db.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
