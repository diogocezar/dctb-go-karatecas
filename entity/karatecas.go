package entity

type Karateca struct {
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

func Create(firstName string, lastName string, birthday string, height float64) *Karateca {
	return &Karateca{
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  birthday,
		Height:    height,
	}
}
