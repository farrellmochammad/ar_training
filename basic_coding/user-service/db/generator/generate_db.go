package main

import (
    "os"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    os.Create("./user_data.db")

    db, err := sql.Open("sqlite3", "./user_data.db")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users ( user_id INTEGER PRIMARY KEY AUTOINCREMENT, user_address VARCHAR(150) NULL, created DATE NULL );")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    db.Close()
}