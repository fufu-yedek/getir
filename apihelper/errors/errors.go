package errors

import (
	"fmt"
	"net/http"
)

const (
	CodeInternalServerError = http.StatusInternalServerError
	CodeUserReadableError   = http.StatusBadRequest
)

var ErrInternalServer = fmt.Errorf("internal server error")
var ErrUserReadable = fmt.Errorf("")

func NewUserReadableErrf(s string, args ...interface{}) error {
	return fmt.Errorf("%w%s", ErrUserReadable, fmt.Sprintf(s, args...))
}
