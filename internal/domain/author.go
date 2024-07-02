package domain

import (
	"net/http"
	"time"
)

type Author struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AuthorRepository interface {
	Create(author Author) error
	GetAll() ([]Author, error)
	GetById(id int) (*Author, error)
	Update(id int, author Author) error
	Delete(id int) error
}

type AuthorService interface {
	Create(author Author) error
	GetAll() ([]Author, error)
	GetById(id int) (*Author, error)
	Update(id int, author Author) error
	Delete(id int) error
}

type AuthorHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
