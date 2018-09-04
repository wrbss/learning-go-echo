package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

type Example struct {
}

func NewExample() (e *Example) {
	return new(Example)
}

func (e Example) GetExample(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok2")
}
