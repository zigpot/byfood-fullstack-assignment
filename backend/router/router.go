package router

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/book-library/handler"
)

func InitRoutes(h *handler.BookHandler) *mux.Router {
	r := mux.NewRouter()

	// Book CRUD
	r.HandleFunc("/books", h.GetBooks).Methods("GET")
	// TODO: Add Create, Update, Delete

	// URL cleanup & redirection
	r.HandleFunc("/process-url", h.ProcessURL).Methods("POST")

	return r
}
