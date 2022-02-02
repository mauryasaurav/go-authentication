package error

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err ApiError) Error() string {
	return err.Message
}

// CustomHTTPErrorHandler implements custom error handler for echo
func CustomHTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message string

	switch he := err.(type) {
	case ApiError:
		code = he.Code
		message = he.Message
		break
	case *echo.HTTPError:
		code = he.Code
		message = he.Message.(string)
		break
	case validator.ValidationErrors:
		code = http.StatusBadRequest
		message = getValidationErrorMessage(he)
		break
	default:
		code = http.StatusInternalServerError
		message = "Something went wrong"
	}

	c.Logger().Error(err)
	_ = c.JSON(code, ApiError{code, message})
}

func getValidationErrorMessage(err validator.ValidationErrors) string {
	return fmt.Sprintf("Invalid value for field `%s`", err[0].Field())
}
