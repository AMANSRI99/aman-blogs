package resp

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data      interface{} `json:"data"`
	Meta      interface{} `json:"meta,omitempty"`
	Message   string      `json:"message,omitempty"`
	ErrorCode string      `json:"error_code,omitempty"`
}

func HTTPServerError(c echo.Context) error {
	r := Response{
		Message:   "Something Went wrong",
		ErrorCode: ServerError,
	}
	return c.JSON(http.StatusInternalServerError, r)
}
