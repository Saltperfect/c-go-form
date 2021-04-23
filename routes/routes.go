package routes

import (
	"fmt"
	"html/template"

	"github.com/gorilla/mux"
	"github.com/saltperfect/c-go-form/controllers"
)

type Router struct {
	multiplexer   *mux.Router
	actionHandler *controllers.ActionHandler
}

func NewRouter(template *template.Template) *Router {
	lshandler := controllers.NewLSHandler()
	uiManager := controllers.NewUIManager(template)
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
	return router.multiplexer
}
