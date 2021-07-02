package entity

import (
	uuid "github.com/satori/go.uuid"
)

type Karateca struct {
	ID        string
	FirstName string
	LastName  string
	Birthday  string
	Height    float64
}

type Karatecas struct {
	ListKaratecas []Karateca
}

func (k *Karatecas) Add(karateca Karateca) {
	k.ListKaratecas = append(k.ListKaratecas, karateca)
}

func Create(id string, firstName string, lastName string, birthday string, height float64) *Karateca {
	var setId = ""
	if id != "" {
		setId = id
	} else {
		setId = uuid.NewV4().String()
	}
	return &Karateca{
		ID:        setId,
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  birthday,
		Height:    height,
	}
}
