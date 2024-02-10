package api

type ApiError struct {
	Description string `json:"description"`
}

func NewApiError(aCode string, aDescription string) *ApiError {
	return &ApiError{
		Description: aDescription,
	}
}
