package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/diogocezar/dctb-go-karatecas/database"
	karatecasModel "github.com/diogocezar/dctb-go-karatecas/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
)

var karatecasMongo = database.Get().Database("karatecas-base").Collection("karatecas")

func GetAll(ctx echo.Context) error {
	cursor, err := karatecasMongo.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var karatecas []bson.M
	if err = cursor.All(context.TODO(), &karatecas); err != nil {
		log.Fatal(err)
	}
	fmt.Println("🤜 List of karatecas returned successfully.")
	fmt.Println(karatecas)
	return ctx.JSON(http.StatusOK, karatecas)
}

func Save(ctx echo.Context) error {
	k := new(karatecasModel.Karateca)
	if err := ctx.Bind(k); err != nil {
		return err
	}
	newKarateca := karatecasModel.Create("", k.FirstName, k.LastName, k.Birthday, k.Height)
	result, err := karatecasMongo.InsertOne(context.TODO(), newKarateca)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("🥋 Karateca was created successfully.")
	fmt.Println(newKarateca)
	return ctx.JSON(http.StatusOK, result)

}

func Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	_, err := karatecasMongo.DeleteOne(context.TODO(), bson.M{"ID": id})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("🥋 Karateca was deleted successfully.")
	fmt.Println(id)
	return ctx.String(http.StatusOK, "DELETED: "+id)
}

func Update(ctx echo.Context) error {
	id := ctx.Param("id")
	k := new(karatecasModel.Karateca)
	if err := ctx.Bind(k); err != nil {
		return err
	}
	newKarateca := karatecasModel.Create(id, k.FirstName, k.LastName, k.Birthday, k.Height)
	_, err := karatecasMongo.UpdateByID(context.TODO(), id, newKarateca)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("🥋 Karateca was updated successfully.")
	fmt.Println(newKarateca)
	return ctx.String(http.StatusOK, "UPDATED: "+id)
}
