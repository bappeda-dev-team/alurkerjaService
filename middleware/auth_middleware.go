package middleware

import (
	"alurkerjaService/model/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Header.Get("X-API-Key") == "RAHASIA" {
			return next(c)
		}

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		return c.JSON(http.StatusUnauthorized, webResponse)
	}
}
