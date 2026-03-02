package app

import (
	"database/sql"
	"os"
	"time"

	"bintangakasyah/belajar-golang-restful-api/helper"
)

func NewDB() *sql.DB {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "root:@tcp(localhost:3306)/db_api"
	}

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}