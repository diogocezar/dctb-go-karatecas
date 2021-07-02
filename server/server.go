package server

import (
	"net/http"

	"github.com/diogocezar/dctb-go-karatecas/data"
	"github.com/diogocezar/dctb-go-karatecas/entity"
	"github.com/labstack/echo/v4"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (server Server) Start() {
	echoServer := echo.New()
	echoServer.GET("/karatecas", server.GetAll)
	echoServer.POST("/karatecas", server.Save)
	echoServer.Logger.Fatal(echoServer.Start(":8888"))
}

func (server Server) GetAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, data.GetAll())
}

func (server Server) Save(ctx echo.Context) error {
	u := new(entity.Karateca)
	if err := ctx.Bind(u); err != nil {
		return err
	}
	newKarateca := entity.Create(u.FirstName, u.LastName, u.Birthday, u.Height)
	data.Save(*newKarateca)
	return ctx.JSON(http.StatusOK, *newKarateca)
}
