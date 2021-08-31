package api

import (
	"github.com/diogocezar/dctb-go-karatecas/handlers"
	"github.com/labstack/echo"
)

func Start(e *echo.Echo) {
	e.GET("/karatecas", handlers.GetAll)
	e.POST("/karatecas", handlers.Save)
	e.DELETE("/karatecas/:id", handlers.Delete)
	e.PUT("/karatecas/:id", handlers.Update)
}
