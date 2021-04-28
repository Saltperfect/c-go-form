package models

import (
	"database/sql"
	"time"

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
    	form VARCHAR(2000),
    	created_at TIMESTAMP
	);`)
	if err != nil {
		panic(err)
	}
}

func (s *SQLite) AddForm( form *Form) error {
	_, err := s.manager.Exec(
		`INSERT INTO FORM (title, form, created_at) VALUES (?, ?, CURRENT_TIMESTAMP )`,
		form.Title, form.Html)
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLite) LoadForm(title string) (*Form, error){
	rows, err := s.manager.Query(`SELECT * from FORM WHERE title =?`, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var form *Form

	rows.Next()
	form = new(Form)
	timestamp := new(time.Time)

	err = rows.Scan(&form.Title, &form.Html, &timestamp)
	if err != nil {
		return nil, err
	}
	form.Created = timestamp.Format("Mon Jan 2 15:04:05 MST 2006")
	
	return form, nil
}

func (s *SQLite) LoadForms() ([]*Form, error){
	rows, err := s.manager.Query(`SELECT * from FORM;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forms []*Form
	var form *Form
	var timestamp time.Time
	var html *string

	for rows.Next(){
		form = new(Form)

		err = rows.Scan(&form.Title, html, &timestamp)

		form.Created = timestamp.Format("Mon Jan 2 15:04:05 MST 2006")
		forms = append(forms, form)
	}
	return forms, nil
}