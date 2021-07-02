package entity

type Karateca struct {
	FirstName string
	LastName  string
	Birthday  string
	Height    float64
}

type Karatecas struct {
	listKaratecas []Karateca
}

func (k *Karatecas) Add(karateca Karateca) {
	k.listKaratecas = append(k.listKaratecas, karateca)
}

// func (k *Karateca) Create(firstName string, lastName string, birthday string, height float64) Karateca {
// 	k.FirstName = firstName
// 	k.LastName = lastName
// 	k.Birthday = birthday
// 	k.Height = height
// 	return k
// }
