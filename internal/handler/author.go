package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/shironxn/go-clean-arch/internal/domain"
)

type AuthorHandler struct {
	service domain.AuthorService
}

func NewAuthorHandler(service domain.AuthorService) domain.AuthorHandler {
	return &AuthorHandler{
		service: service,
	}
}

func (h *AuthorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Author

	reqBody, _ := io.ReadAll(r.Body)
	_ = json.Unmarshal(reqBody, &req)

	err := h.service.Create(req)
	if err != nil {
		res := domain.ErrorResponse{
			Error: err.Error(),
		}
		response, _ := json.Marshal(res)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	}

	res := domain.SuccessRespons{
		Message: "Successfully create author",
		Data:    nil,
	}
	response, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(response)
}

func (h *AuthorHandler) GetAll(w http.ResponseWriter, r *http.Request) {}

func (h *AuthorHandler) GetById(w http.ResponseWriter, r *http.Request) {}

func (h *AuthorHandler) Update(w http.ResponseWriter, r *http.Request) {}

func (h *AuthorHandler) Delete(w http.ResponseWriter, r *http.Request) {}
