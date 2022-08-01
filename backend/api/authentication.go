package api

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

func getKey(c echo.Context) error {
	return c.String(http.StatusOK, "Getting key")
}

func login(c echo.Context) error {
	return c.String(http.StatusOK, "Accessable")
}

func createNewUser(c echo.Context) error {
	var data map[string]string

	err := c.Bind(&data)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}
