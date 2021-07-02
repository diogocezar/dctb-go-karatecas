package main

import (
	"fmt"

	"github.com/diogocezar/dctb-go-karatecas/entity"
)

func main() {
	k1 := entity.Karateca{
		FirstName: "Diogo",
		LastName:  "Cezar",
		Birthday:  "19/02/1986",
		Height:    1.69,
	}
	list := entity.Karatecas{}
	list.Add(k1)
	fmt.Println(list)
}
