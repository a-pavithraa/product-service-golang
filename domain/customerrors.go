package domain

type ProductNameFoundError struct {
	message string
}

func (e *ProductNameFoundError) Error() string {
	return e.message
}

func NewProductNameFoundError(message string) *ProductNameFoundError {
	return &ProductNameFoundError{
		message: message,
	}
}

type ProductNotFoundError struct {
	message string
}

func (e *ProductNotFoundError) Error() string {
	return e.message
}
func NewProductNotFoundError(message string) *ProductNotFoundError {
	return &ProductNotFoundError{
		message: message,
	}
}
