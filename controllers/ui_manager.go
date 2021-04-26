package controllers

import (
	"html/template"
	"net/http"
)

type UIManager struct {
	template *template.Template
}

func NewUIManager(template *template.Template) *UIManager {
	return &UIManager{
		template: template,
	}
}

func (ui *UIManager) RenderPage(w http.ResponseWriter, tmpl string, page interface{}) {
	err := ui.template.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
