package exceptions

type BadRequestError struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e BadRequestError) Error() string {
	return e.Message
}

func (e BadRequestError) StatusCode() int {
	return 400
}

type ValidationError = BadRequestError

type UnprocessableError struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors"`
}

func (e UnprocessableError) Error() string {
	return e.Message
}

func (e UnprocessableError) StatusCode() int {
	return 422
}
