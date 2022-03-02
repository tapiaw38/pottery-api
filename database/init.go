package database

import (
	"database/sql"
	"log"
	"sync"
)

var (
	data *sql.DB
	once sync.Once
)

// InitDB is the function that initializes the database
func InitDB() *sql.DB {

	db, err := Connect()

	if err != nil {
		panic(err)
	}

	err = MakeMigration(db)

	if err != nil {
		panic(err)
	}

	log.Println("Migration complete")

	return db
}

// NewConnection is the function to get a only connection
func NewConnection() *sql.DB {

	once.Do(func() {
		data = InitDB()
	})

	return data
}

// Check connection to the database
func CheckConnection() bool {

	err := data.Ping()

	if err != nil {
		return false
	}

	return true
}
