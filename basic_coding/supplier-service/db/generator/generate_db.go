package main

import (
    "os"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    os.Create("./supplier_data.db")

    db, err := sql.Open("sqlite3", "./supplier_data.db")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS supplier_seller ( supplier_id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(64) NULL, created DATE NULL );")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS supplier_store ( store_id INTEGER PRIMARY KEY AUTOINCREMENT, supplier_id INTEGER NULL, store_name VARCHAR(64) NULL, store_address VARCHAR(150) NULL, created DATE NULL );")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	_,err = db.Exec("CREATE TABLE IF NOT EXISTS product ( product_id INTEGER PRIMARY KEY AUTOINCREMENT,user_id INTEGER NULL, product_name VARCHAR(64) NULL, product_stock INTEGER NULL, created DATE NULL );")
	
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }


    db.Close()
}