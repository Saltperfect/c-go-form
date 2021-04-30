package routes

import (
	"fmt"
	"html/template"

	"github.com/gorilla/mux"
	"github.com/saltperfect/c-go-form/controllers"
	"github.com/saltperfect/c-go-form/models"
)

type Router struct {
	multiplexer   *mux.Router
	actionHandler *controllers.ActionHandler
}

func NewRouter(db models.Database, templates *template.Template) *Router {
	lshandler := controllers.NewLSHandler(templates, db)
	uiManager := controllers.NewUIManager(templates)
	return &Router{
		multiplexer:   mux.NewRouter(),
		actionHandler: controllers.NewActionHandler(lshandler, uiManager),
	}
}

func (router *Router) GetRoutes() *mux.Router {
	fmt.Println("enter into getroutes")
	router.multiplexer.HandleFunc("/view/{title}", router.actionHandler.View)
	router.multiplexer.HandleFunc("/edit/{title}", router.actionHandler.Edit)
	router.multiplexer.HandleFunc("/save/{title}", router.actionHandler.Save)
	router.multiplexer.HandleFunc("/create/", router.actionHandler.Create)
	router.multiplexer.HandleFunc("/generate/", router.actionHandler.Generate)
	router.multiplexer.HandleFunc("/list/", router.actionHandler.List)
	router.multiplexer.HandleFunc("/viewform/{title}", router.actionHandler.ViewForm)
	router.multiplexer.HandleFunc("/submit/{title}", router.actionHandler.SubmitForm)
	return router.multiplexer
}
