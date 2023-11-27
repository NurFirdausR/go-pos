package database

import (
	"database/sql"
	"time"

	"github.com/NurFirdausR/go-pos/helper"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Bismillah@123@tcp(127.0.0.1:3306)/golang_pos?parseTime=true")
	helper.PanicIfError(err)

	// Test the connection
	err = db.Ping()
	if err != nil {
		helper.PanicIfError(err)
		// fmt.Println("Error:", err)
		// return
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db

}
