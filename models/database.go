package models

type Database interface {
	AddForm( *Form) error
	LoadForm( string ) (*Form, error)
	LoadForms() ([]*Form, error)
}