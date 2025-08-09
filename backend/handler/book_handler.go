package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yourusername/book-library/repository"
)

type BookHandler struct {
	Repo *repository.BookRepo
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Repo.GetAll()
	if err != nil {
		http.Error(w, "Failed to fetch books", 500)
		log.Println("Error fetching books:", err)

		return
	}
	json.NewEncoder(w).Encode(books)
}
