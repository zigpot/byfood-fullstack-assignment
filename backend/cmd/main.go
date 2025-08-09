package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/book-library/config"
	"github.com/yourusername/book-library/db"
	"github.com/yourusername/book-library/handler"
	"github.com/yourusername/book-library/repository"
	"github.com/yourusername/book-library/router"
	"os"
)

func main() {
	wd, _ := os.Getwd()

	fmt.Println("Working dir:", wd)
	config.LoadEnv()

	conn, err := db.Connect()
	if err != nil {
		log.Fatal("DB connection error: ", err)
	}

	repo := &repository.BookRepo{DB: conn}
	h := &handler.BookHandler{Repo: repo}
	r := router.InitRoutes(h)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
