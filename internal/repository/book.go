package repository

import (
	"database/sql"

	"github.com/shironxn/go-clean-arch/internal/domain"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) domain.BookRepository {
	return &BookRepository{
		DB: db,
	}
}

func (r *BookRepository) Create(entity domain.Book) (*domain.Book, error) {
	q := "INSERT INTO books (title, author, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	res, err := r.DB.Exec(q)
	if err != nil {
		return nil, err
	}
	res.LastInsertId()

	return nil, nil
}

func (r *BookRepository) GetAll() ([]domain.Book, error) {
	return nil, nil
}

func (r *BookRepository) GetById(id int) (*domain.Book, error) {
	return nil, nil
}

func (r *BookRepository) Update(entity domain.Book, entityUpdate domain.Book) (*domain.Book, error) {
	return nil, nil
}

func (r *BookRepository) Delete(id int) error {
	return nil
}
