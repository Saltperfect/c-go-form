package routes

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/saltperfect/c-go-form/controllers"
)

type Router struct {
	multiplexer   *mux.Router
	actionHandler *controllers.ActionHandler
}

func NewRouter() *Router {
	return &Router{
		multiplexer:   mux.NewRouter(),
		actionHandler: controllers.NewActionHandler(),
	}
}

func (router *Router) GetRoutes() *mux.Router {
	fmt.Println("enter into getroutes")
	router.multiplexer.HandleFunc("/view/", router.actionHandler.View).Methods("POST")
	return router.multiplexer
}
