package routes

import (
	"github.com/gorilla/mux"
	"github.com/harshitasao/BookStore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/", controllers.DeleteBook).Methods("DELETE")
}
