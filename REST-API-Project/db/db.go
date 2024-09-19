package db

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

// global variable for database
type DB struct {
	*sql.DB
}

var DB_conn *DB

func InitDB() {
	var err error
	DB_conn, err = NewDB("./api.db")

	if err != nil {
		panic("Could not create database")
	}

	createTables(DB_conn)
}

func createTables(DB_conn *DB) {

	createUsersTables := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB_conn.Exec(createUsersTables)

	if err != nil {
		panic("could not create users table.")
	}

	createEventsTables := `
			CREATE TABLE IF NOT EXISTS events (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER, 
			FOREIGN KEY(user_id) REFERENCES users(id)
			);`
	_, err = DB_conn.ExecContext(context.Background(), createEventsTables)

	if err != nil {
		panic("could not create events table.")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB_conn.Exec(createRegistrationsTable)

	if err != nil {
		panic(err)
	}
}

func NewDB(datasourceName string) (*DB, error) {
	DB_conn, _ := sql.Open("sqlite", datasourceName)

	DB_conn.SetMaxOpenConns(10)
	DB_conn.SetMaxIdleConns(5)

	return &DB{DB_conn}, nil

}
