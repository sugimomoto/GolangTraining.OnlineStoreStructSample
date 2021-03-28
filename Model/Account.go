package Model

type Account struct {
	FirstName string
	LastName  string
}

func (acc *Account) ChangeName(firstName string, lastName string) {
	acc.FirstName = firstName
	acc.LastName = lastName
}
