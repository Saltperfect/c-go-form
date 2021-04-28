package controllers

import (
	"io/ioutil"
	"text/template"
	"bytes"
	"github.com/saltperfect/c-go-form/models"
)

type LSHandler struct {
	template *template.Template
	db models.Database
}

func NewLSHandler(template *template.Template, db models.Database) *LSHandler {
	return &LSHandler{
		template: template,
		db: db,
	}
}

func (ls *LSHandler) Save(p *models.Page) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("data/"+filename, p.Body, 0600)
}

func (ls *LSHandler) LoadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("data/" + filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}

func (ls *LSHandler) SaveHtml(title string, tmpl string, page interface{}) error {

	var buf bytes.Buffer
	err := ls.template.ExecuteTemplate(&buf, tmpl, page)
	if err != nil {
		return err
	}
	data := &models.Form{
		Title: title,
		Html: buf.String(),
	}
	ls.db.AddForm(data)
	return nil
}
