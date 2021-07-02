package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct{}

func (server Server) start() {
	echoServer := echo.New()
	echoServer.GET("/karatecas", server.getAll)
	echoServer.Logger.Fatal(echoServer.Start(":8888"))
}

func (server Server) getAll(ctx) {}
