package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// NotFoundHandler – аналог notFound middleware из Express
func NotFoundHandler(c echo.Context) error {
	return c.JSON(http.StatusNotFound, map[string]string{
		"message": "Not found - " + c.Request().URL.Path,
	})
}

// ErrorHandler – аналог errorHandler из Express
func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := err.Error()
	stack := errors.New("nill")

	// Проверка, является ли ошибка HTTPError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		if str, ok := he.Message.(string); ok {
			msg = str
		} else {
			msg = http.StatusText(code)
		}
	}

	// Если окружение не продакшен – логгируем стек
	if os.Getenv("NODE_ENV") != "production" {
		log.Printf("ERROR: %+v\n", err)
		stack = err
	}

	// Ответ
	_ = c.JSON(code, map[string]interface{}{
		"message": msg,
		"stack":   stack,
	})
}
