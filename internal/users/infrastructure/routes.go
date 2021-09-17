package infrastructure

import (
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, h Handler) {
	e.POST("/users", h.Create)
	e.PUT("/users/:id", h.Update)
	e.GET("/users/:id", h.FindById)
	e.GET("/users", h.FindByDocument)
}
