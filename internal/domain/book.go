package domain

import (
	"net/http"
	"time"
)

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	AuthorId  int       `json:"author_id"`
	Genre     string    `json:"genre"`
	Synopsis  string    `json:"synopsis"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookRepository interface {
	Create(book Book) error
	GetAll() ([]Book, error)
	GetById(id int) (*Book, error)
	Update(id int, book Book) error
	Delete(id int) error
}

type BookService interface {
	Create(book Book) error
	GetAll() ([]Book, error)
	GetById(id int) (*Book, error)
	Update(id int, book Book) error
	Delete(id int) error
}

type BookHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
