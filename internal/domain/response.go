package domain

type SuccessRespons struct {
	Message string
	Data    any
}

type ErrorResponse struct {
	Error string
}
