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
	echoServer.DELETE("/karatecas/:id", server.Delete)
	echoServer.PUT("/karatecas/:id", server.Update)
	echoServer.Logger.Fatal(echoServer.Start(":8888"))
}

func (server Server) GetAll(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, data.GetAll())
}

func (server Server) Save(ctx echo.Context) error {
	k := new(entity.Karateca)
	if err := ctx.Bind(k); err != nil {
		return err
	}
	newKarateca := entity.Create("", k.FirstName, k.LastName, k.Birthday, k.Height)
	data.Save(*newKarateca)
	return ctx.JSON(http.StatusOK, *newKarateca)
}

func (server Server) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	data.Delete(id)
	return ctx.String(http.StatusOK, "DELETED: "+id)
}

func (server Server) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	k := new(entity.Karateca)
	if err := ctx.Bind(k); err != nil {
		return err
	}
	newKarateca := entity.Create(id, k.FirstName, k.LastName, k.Birthday, k.Height)
	data.Update(*newKarateca, id)
	return ctx.String(http.StatusOK, "UPDATED: "+id)
}
