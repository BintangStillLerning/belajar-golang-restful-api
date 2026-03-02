package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(Error string) NotFoundError {
	return NotFoundError{Error: Error}
}