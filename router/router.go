package router

import (
	"github.com/diogocezar/dctb-go-karatecas/api"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	api.Start(e)
	return e
}
