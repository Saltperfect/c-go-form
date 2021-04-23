package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/davecgh/go-spew/spew"
	"github.com/saltperfect/c-go-form/routes"
)

func main() {
	// templates := template.Must(template.ParseGlob("./template/*"))
	routers := routes.NewRouter()

	mux := routers.GetRoutes()
	mux.HandleFunc("/he", func(http.ResponseWriter, *http.Request) {
		fmt.Println("called he")
	})
	spew.Dump(routers)
	log.Fatal(http.ListenAndServe(":8080", mux))

}
