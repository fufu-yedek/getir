package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	CodeInternalServerError = http.StatusInternalServerError
	CodeUserReadableError   = http.StatusBadRequest
)

var ErrInternalServer = fmt.Errorf("internal server error")
var ErrUserReadable = errors.New("")

func NewErrUserReadable(s string, args ...interface{}) error {
	return fmt.Errorf("%w%s", ErrUserReadable, fmt.Sprintf(s, args...))
}
