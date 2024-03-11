package errors

type ErrInvalidID struct {
	Description string `json:"description"`
}

type ErrInvalidBody struct {
	Description string `json:"description"`
}

func (e *ErrInvalidID) Error() string {
	return e.Description
}

func (e *ErrInvalidBody) Error() string {
	return e.Description
}
