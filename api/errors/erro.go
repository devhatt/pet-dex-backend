package errors

type ErrInvalidID struct {
	Description string `json:"description"`
}

type InvalidBody struct {
	Description string `json:"description"`
}
