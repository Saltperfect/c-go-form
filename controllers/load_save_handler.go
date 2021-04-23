package controllers

import (
	"io/ioutil"

	"github.com/saltperfect/c-go-form/models"
)

type LSHandler struct {
}

func NewLSHandler() *LSHandler {
	return new(LSHandler)
}

func (*LSHandler) Save(p *models.Page) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("data/"+filename, p.Body, 0600)
}

func (*LSHandler) LoadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("data/" + filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}
