package models

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func NewDirector(firstname string, lastname string) *Director {
	return &Director{
		FirstName: firstname,
		LastName:  lastname,
	}
}
