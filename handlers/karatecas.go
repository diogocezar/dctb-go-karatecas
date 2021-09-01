package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/diogocezar/dctb-go-karatecas/database"
	karatecasModel "github.com/diogocezar/dctb-go-karatecas/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var karatecasMongo = database.Get().Database("karatecas-base").Collection("karatecas")

func GetAll(ctx echo.Context) error {
	cursor, err := karatecasMongo.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	var karatecas []bson.M
	if err = cursor.All(context.TODO(), &karatecas); err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	fmt.Println("🤜 List of karatecas returned successfully.")
	fmt.Println(karatecas)
	return ctx.JSON(http.StatusOK, karatecas)
}

func GetOne(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	cursor, err := karatecasMongo.Find(context.TODO(), bson.M{"_id": id})
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	var karatecas []bson.M
	if err = cursor.All(context.TODO(), &karatecas); err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	fmt.Println("🤜 Finded karatecas returned successfully.")
	fmt.Println(karatecas)
	return ctx.JSON(http.StatusOK, karatecas)
}

func Save(ctx echo.Context) error {
	k := new(karatecasModel.Karateca)
	if err := ctx.Bind(k); err != nil {
		return err
	}
	newKarateca := karatecasModel.Create(k.FirstName, k.LastName, k.Birthday, k.Height)
	result, err := karatecasMongo.InsertOne(context.TODO(), newKarateca)
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	fmt.Println("🥋 Karateca was created successfully.")
	fmt.Println(newKarateca)
	return ctx.JSON(http.StatusOK, result)

}

func Delete(ctx echo.Context) error {
	id, errConvert := primitive.ObjectIDFromHex(ctx.Param("id"))
	if errConvert != nil {
		fmt.Println(errConvert)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	_, err := karatecasMongo.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	fmt.Println("🥋 Karateca was deleted successfully.")
	fmt.Println(id)
	return ctx.JSON(http.StatusOK, "🥋 Karateca with id "+ctx.Param("id")+"was deleted successfully.")
}

func Update(ctx echo.Context) error {
	id, errConvert := primitive.ObjectIDFromHex(ctx.Param("id"))
	if errConvert != nil {
		fmt.Println(errConvert)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	fmt.Println(id)
	k := new(karatecasModel.Karateca)
	if err := ctx.Bind(k); err != nil {
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{
		"firstname": k.FirstName,
		"lastname":  k.LastName,
		"birthday":  k.Birthday,
		"height":    k.Height,
	}}

	result, err := karatecasMongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return ctx.String(http.StatusInternalServerError, "Internal server error.")
	}
	fmt.Println("🥋 Karateca was updated successfully.")
	return ctx.JSON(http.StatusOK, result)
}
