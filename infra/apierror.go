package api

type ApiError struct {
	description string
}

func NewApiError(aCode string, aDescription string) *ApiError {
	return &ApiError{
		description: aDescription,
	}
}
