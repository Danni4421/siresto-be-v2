package exceptions

type BaseError interface {
	Error() string
	StatusCode() int
}
