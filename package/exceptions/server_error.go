package exceptions

type ServerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ServerError) Error() string {
	return e.Message
}

func (e *ServerError) StatusCode() int {
	return e.Code
}

const (
	StatusInternalServerError           = 500
	StatusNotImplemented                = 501
	StatusBadGateway                    = 502
	StatusServiceUnavailable            = 503
	StatusGatewayTimeout                = 504
	StatusHTTPVersionNotSupported       = 505
	StatusNetworkAuthenticationRequired = 511
)

func NewInternalServerError(message string) *ServerError {
	return &ServerError{Code: StatusInternalServerError, Message: message}
}

func NewNotImplemented(message string) *ServerError {
	return &ServerError{Code: StatusNotImplemented, Message: message}
}

func NewBadGateway(message string) *ServerError {
	return &ServerError{Code: StatusBadGateway, Message: message}
}

func NewServiceUnavailable(message string) *ServerError {
	return &ServerError{Code: StatusServiceUnavailable, Message: message}
}

func NewGatewayTimeout(message string) *ServerError {
	return &ServerError{Code: StatusGatewayTimeout, Message: message}
}

func NewHTTPVersionNotSupported(message string) *ServerError {
	return &ServerError{Code: StatusHTTPVersionNotSupported, Message: message}
}

func NewNetworkAuthenticationRequired(message string) *ServerError {
	return &ServerError{Code: StatusNetworkAuthenticationRequired, Message: message}
}
