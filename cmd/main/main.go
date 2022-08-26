package main

import (
	"log"
	"net/http"

	"github.com/fikma/go-pustakaBuku/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterPustakaBukuRoutes(r)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
