package models

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type SQLite struct {
	manager *sql.DB
}

var sqliteInstance *SQLite = nil

func NewSQLiteDB() *SQLite {
	if sqliteInstance == nil {
		db, err := sql.Open("sqlite", "./form.db")
		if err != nil {
			panic(err)
		}
		mustInitSQLiteDB(db)
		sqliteInstance = &SQLite{manager: db}
	}
	return sqliteInstance
}

func mustInitSQLiteDB( db *sql.DB){
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS 
    FORM (
    	title VARCHAR(255) PRIMARY KEY , 
    	form VARCHAR(2000000),
    	created_at TIMESTAMP,
	);`)
	if err != nil {
		panic(err)
	}
}

func (s *SQLite) AddForm( form *Form) {
	
}