package controllers

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/saltperfect/c-go-form/models"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type LSHandler struct {
	template *template.Template
	db       models.Database
}

func NewLSHandler(template *template.Template, db models.Database) *LSHandler {
	return &LSHandler{
		template: template,
		db:       db,
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
		Html:  buf.String(),
	}
	err = ls.db.AddForm(data)
	if err != nil {
		return err
	}
	return nil
}

func (ls *LSHandler) LoadForm(title string) (string, error) {
	form, err := ls.db.LoadForm(title)
	spew.Dump(title, form)
	formioreader := strings.NewReader(form.Html)
	node, err := html.Parse(formioreader)
	if err != nil {
		return "", err
	}

	var f func(*html.Node)
	var submitButton = &html.Node{
		Type:     html.ElementNode,
		DataAtom: atom.Input,
		Data:     "input",
		Attr: []html.Attribute{
			{
				Key: "type",
				Val: "submit",
			},
			{
				Key: "value",
				Val: "Submit Form",
			},
			{
				Key: "formaction",
				Val: "/submit/" + title,
			},
		},
	}
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "form" {
			n.AppendChild(submitButton)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(node)
	var b bytes.Buffer
	err = html.Render(&b, node)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
func (ls *LSHandler) LoadForms() ([]*models.Form, error) {
	forms, err := ls.db.LoadForms()
	if err != nil {
		return nil, err
	}
	return forms, nil
}
