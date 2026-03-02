package app

import (
	"database/sql"
	"log"
	"os"
	"time"
)

func NewDB() *sql.DB {

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Println("DATABASE_URL not set")
		return nil
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("DB OPEN ERROR:", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println("DB PING ERROR:", err)
		return nil
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	log.Println("DATABASE CONNECTED")


	return db
}