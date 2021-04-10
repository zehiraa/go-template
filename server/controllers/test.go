package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTest(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}