package domain

import (
	"net/http"
	"time"
)

type Book struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    Author    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookRepository interface {
	Create(entity Book) (*Book, error)
	GetAll() ([]Book, error)
	GetById(id int) (*Book, error)
	Update(entity Book, entityUpdate Book) (*Book, error)
	Delete(id int) error
}

type BookService interface {
	Create(entity Book) (*Book, error)
	GetAll() ([]Book, error)
	GetById(id int) (Book, error)
	Update(entity Book, entityUpdate Book) (*Book, error)
	Delete(id int) error
}

type BookHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
