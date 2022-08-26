package routes

import (
	"github.com/fikma/go-pustakaBuku/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterPustakaBukuRoutes = func(router *mux.Router) {
	router.HandleFunc("/buku/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/buku/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/buku/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/buku/{bookId}", controllers.UpdateBookById).Methods("PUT")
	router.HandleFunc("/buku/{bookId}", controllers.DeleteBookById).Methods("DELETE")
}
