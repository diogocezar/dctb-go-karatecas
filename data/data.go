package data

import "github.com/diogocezar/dctb-go-karatecas/entity"

var Karatecas entity.Karatecas

func LoadAll() {
	k1 := entity.Create("Diogo", "Cezar", "19/02/1986", 1.69)
	k2 := entity.Create("Diogo", "Cezar", "19/02/1986", 1.69)
	Karatecas.Add(*k1)
	Karatecas.Add(*k2)
}

func GetAll() []entity.Karateca {
	return Karatecas.ListKaratecas
}

func Save(karateca entity.Karateca) {
	Karatecas.Add(karateca)
}
