package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/silverhand7/go-book-rest-api/pkg/controllers"
)

var RegisterBookRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{bookId}", controllers.GetBookById).Methods(http.MethodGet)
	router.HandleFunc("/books/{bookId}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books", controllers.DeleteBook).Methods(http.MethodDelete)
}
