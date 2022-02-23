package message_tool

import (
	"errors"
	"net/http"

	"github.com/samuskitchen/go-message-tool/errors"

	"github.com/labstack/echo/v4"
	"github.com/samuskitchen/go-context-tool"
	"github.com/sirupsen/logrus"
)

const (
	responseData    = "data"
	responseMessage = "info"
	responseError   = "error"
)

type ResponseDetails struct {
	NumItems int    `json:"num_items"`
	OffSet   int    `json:"offset"`
	Limit    int    `json:"limit,omitempty"`
	Omits    string `json:"omits,omitempty"`
}

type Response struct {
	Type    string           `json:"type,omitempty"` //error, response
	Message string           `json:"message,omitempty"`
	Data    interface{}      `json:"data,omitempty"`
	Details *ResponseDetails `json:"details,omitempty"`
}

func Success(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusOK, &Response{
		Type: responseData,
		Data: payload,
	})
}

func SuccessDetails(c echo.Context, ctx context_tool.ContextToolInterface, payload interface{}, numElements int) error {
	params := ctx.GetParams()
	return c.JSON(http.StatusOK, &Response{
		Type: responseData,
		Data: payload,
		Details: &ResponseDetails{
			OffSet:   params.OffSet(),
			Limit:    params.Limit(),
			Omits:    params.SkipFields(),
			NumItems: numElements,
		},
	})
}

func Message(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, &Response{
		Type:    responseMessage,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, err error) error {
	var customError *errores.CustomError
	var code int = 400
	var message string = "something happened, there was an unexpected error"

	if errors.As(err, &customError) {
		code = customError.HttpCode
		message = customError.ErrorMessage
	}

	go func(e error, ec *errores.CustomError) {
		if ec == nil {
			logrus.Error(e.Error())
			return
		}

		if ec.GetError() != nil {
			logrus.Error(e.Error())
			return
		}
	}(err, customError)

	return c.JSON(code, &Response{Type: responseError, Message: message})
}

func JSONErrorResponse(c echo.Context) error {
	return ErrorResponse(c, errores.NewBadRequest(nil, errores.ErrInvalidJSON))
}

func QueryErrorResponse(c echo.Context) error {
	return ErrorResponse(c, errores.NewBadRequest(nil, errores.ErrInvalidQueryParam))
}
