package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/silverhand7/go-book-rest-api/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
