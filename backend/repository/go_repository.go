package repository

import (
	"database/sql"

	"github.com/yourusername/book-library/models"
)

type BookRepo struct {
	DB *sql.DB
}

func (r *BookRepo) GetAll() ([]models.Book, error) {
	rows, err := r.DB.Query(`
    SELECT 
        id, 
        title, 
        author, 
        COALESCE(year, EXTRACT(YEAR FROM published_date)::int, 0) AS year 
    FROM books
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// Add Create, Update, Delete similarly...
