package exceptions

type ClientError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ClientError) Error() string {
	return e.Message
}

func (e ClientError) StatusCode() int {
	return e.Code
}

const (
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusConflict            = 409
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429
)

func NewBadRequest(message string) ClientError {
	return ClientError{Code: StatusBadRequest, Message: message}
}

func NewUnauthorized(message string) ClientError {
	return ClientError{Code: StatusUnauthorized, Message: message}
}

func NewForbidden(message string) ClientError {
	return ClientError{Code: StatusForbidden, Message: message}
}

func NewNotFound(message string) ClientError {
	return ClientError{Code: StatusNotFound, Message: message}
}

func NewMethodNotAllowed(message string) ClientError {
	return ClientError{Code: StatusMethodNotAllowed, Message: message}
}

func NewConflict(message string) ClientError {
	return ClientError{Code: StatusConflict, Message: message}
}

func NewUnprocessableEntity(message string) ClientError {
	return ClientError{Code: StatusUnprocessableEntity, Message: message}
}

func NewTooManyRequests(message string) ClientError {
	return ClientError{Code: StatusTooManyRequests, Message: message}
}
