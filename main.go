package main

import (
	"log"
	"net/http"
	"html/template"
	"github.com/saltperfect/c-go-form/routes"
)

func main() {
	templates := template.Must(template.ParseGlob("./template/*"))
	routers := routes.NewRouter(templates)

	mux := routers.GetRoutes()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
