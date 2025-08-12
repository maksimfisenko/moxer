package errorsx

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrEmailAlreadyExists = errors.New("user with this email already exists")
	ErrInvalidUserId      = errors.New("invalid user id used in template")
)

var (
	ErrUnauthorizedHTTP       = echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid auth token")
	ErrInvalidRequestBodyHTTP = echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	ErrUserExistsHTTP         = echo.NewHTTPError(http.StatusConflict, "user already exists")
	ErrInternalServerHTTP     = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	ErrUserNotFoundHTTP       = echo.NewHTTPError(http.StatusBadRequest, "user not found")
	ErrInvalidTokenHTTP       = echo.NewHTTPError(http.StatusBadRequest, "invalid authentication token")
	ErrInvalidTemplateIdHTTP  = echo.NewHTTPError(http.StatusBadRequest, "invalid template id parameter")
	ErrTemplateNotFoundHTTP   = echo.NewHTTPError(http.StatusBadRequest, "template not found")
	ErrTemplateExistsHTTP     = echo.NewHTTPError(http.StatusBadRequest, "template with given name already exists")
)

type Errorx struct {
	Code    string
	Message string
	Err     error
}

func (e *Errorx) Error() string {
	return e.Message
}

func (e *Errorx) Unwrap() error {
	return e.Err
}

func New(code, msg string, err error) *Errorx {
	return &Errorx{
		Code:    code,
		Message: msg,
		Err:     err,
	}
}

func Is(err error, code string) bool {
	var errorx *Errorx
	if errors.As(err, &errorx) {
		return errorx.Code == code
	}
	return false
}
