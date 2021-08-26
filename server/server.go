package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/diogocezar/dctb-go-karatecas/dbmongo"
	"github.com/diogocezar/dctb-go-karatecas/entity"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type Server struct{}

var karatecasMongo = dbmongo.Get().Database("karatecas-base").Collection("karatecas")

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
	cursor, err := karatecasMongo.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var karatecas []bson.M
	if err = cursor.All(context.TODO(), &karatecas); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, karatecas)
}

func (server Server) Save(ctx echo.Context) error {
	k := new(entity.Karateca)
	if err := ctx.Bind(k); err != nil {
		return err
	}
	newKarateca := entity.Create("", k.FirstName, k.LastName, k.Birthday, k.Height)
	result, err := karatecasMongo.InsertOne(context.TODO(), newKarateca)
	if err != nil {
		fmt.Println(err)
	}
	// data.Save(*newKarateca)
	return ctx.JSON(http.StatusOK, result)

}

func (server Server) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	_, err := karatecasMongo.DeleteOne(context.TODO(), bson.M{"ID": id})
	if err != nil {
		fmt.Println(err)
	}
	//data.Delete(id)
	return ctx.String(http.StatusOK, "DELETED: "+id)
}

func (server Server) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	k := new(entity.Karateca)
	if err := ctx.Bind(k); err != nil {
		return err
	}
	newKarateca := entity.Create(id, k.FirstName, k.LastName, k.Birthday, k.Height)
	_, err := karatecasMongo.UpdateByID(context.TODO(), id, newKarateca)
	if err != nil {
		fmt.Println(err)
	}
	//data.Update(*newKarateca, id)
	return ctx.String(http.StatusOK, "UPDATED: "+id)
}
