package data

import (
	"fmt"

	"github.com/diogocezar/dctb-go-karatecas/entity"
)

var Karatecas entity.Karatecas

func LoadAll() {
	k1 := entity.Create("", "Diogo", "Cezar", "19/02/1986", 1.69)
	k2 := entity.Create("", "Diogo", "Cezar", "19/02/1986", 1.69)
	Karatecas.Add(*k1)
	Karatecas.Add(*k2)
}

func GetAll() []entity.Karateca {
	return Karatecas.ListKaratecas
}

func SetAll(listKaratecas []entity.Karateca) {
	Karatecas.ListKaratecas = listKaratecas
}

func Save(karateca entity.Karateca) {
	Karatecas.Add(karateca)
}

func Delete(id string) {
	fmt.Println(id)
	SetAll(deleteElement(Karatecas.ListKaratecas, id))
}

func Update(k entity.Karateca, id string) {
	Delete(id)
	Save(k)
}

func deleteElement(karatecas []entity.Karateca, id string) []entity.Karateca {
	index := linearSearch(karatecas, id)
	if index != -1 {
		return append(karatecas[:index], karatecas[index+1:]...)
	} else {
		return karatecas
	}
}

func linearSearch(karatecas []entity.Karateca, id string) int {
	for i, n := range karatecas {
		if n.ID == id {
			return i
		}
	}
	return -1
}
