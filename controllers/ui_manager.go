package controllers

import (
	"html/template"
	"net/http"

	"github.com/saltperfect/c-go-form/models"
)

type UIManager struct {
	template *template.Template
}

func NewUIManager(template *template.Template) *UIManager {
	return &UIManager{
		template: template,
	}
}

func (ui *UIManager) Render(w http.ResponseWriter, tmpl string, page *models.Page) {
	err := ui.template.ExecuteTemplate(w, tmpl+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
