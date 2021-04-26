package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
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
	ah.uiManager.RenderPage(w, "view", p)
}

func (ah *ActionHandler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p, err := ah.lshandler.LoadPage(vars["title"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	ah.uiManager.RenderPage(w, "edit", p)
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

func (ah *ActionHandler) Generate(w http.ResponseWriter, r *http.Request) {
	label := r.FormValue("jsondata")
	var list models.ElementList
	fmt.Println(label)
	err := json.Unmarshal([]byte(label), &list)
	if err != nil {
		ah.uiManager.RenderPage(w, "error", err.Error())
	}
	spew.Dump(list)
	ah.uiManager.RenderPage(w, "input", list.List)
}

func (ah *ActionHandler) Create(w http.ResponseWriter, r *http.Request) {
	ah.uiManager.RenderPage(w, "create", &models.Page{})
}
