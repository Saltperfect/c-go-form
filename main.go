package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/saltperfect/c-go-form/models"
	"github.com/saltperfect/c-go-form/routes"
)

func main() {
	templates := template.Must(template.ParseGlob("./template/*"))
	dbmanager := models.NewSQLiteDB()
	routers := routes.NewRouter(dbmanager, templates)

	mux := routers.GetRoutes()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
