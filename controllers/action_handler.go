package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/saltperfect/c-go-form/models"
)

type ActionHandler struct {
	lshandler *LSHandler
	uiManager *UIManager
}

func NewActionHandler(lshandler *LSHandler, uiManager *UIManager) *ActionHandler {
	return &ActionHandler{
		lshandler: lshandler,
		uiManager: uiManager,
	}
}

func (ah *ActionHandler) View(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := ah.lshandler.LoadPage(vars["title"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	ah.uiManager.Render(w, "view", p)
}

func (ah *ActionHandler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := ah.lshandler.LoadPage(vars["title"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	ah.uiManager.Render(w, "edit", p)
}

func (ah *ActionHandler) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["title"]
	body := r.FormValue("body")
	p := &models.Page{Title: t, Body: []byte(body)}
	err := ah.lshandler.Save(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+t, http.StatusFound)
}
