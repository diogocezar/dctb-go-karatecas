package karatecasModel

type Karateca struct {
	FirstName string
	LastName  string
	Birthday  string
	Height    float64
}

func Create(firstName string, lastName string, birthday string, height float64) *Karateca {
	return &Karateca{
		FirstName: firstName,
		LastName:  lastName,
		Birthday:  birthday,
		Height:    height,
	}
}
