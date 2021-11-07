package utils

type AppError struct {
	Message string `json:"message"`
	Code    uint   `json:"code,omitempty"`
}

func (e AppError) Error() string {
	return e.Message
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}
