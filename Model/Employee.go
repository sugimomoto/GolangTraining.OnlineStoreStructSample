package Model

type Employee struct {
	Account

	Credit float64
}

func (employee *Employee) AddCredits(credit float64) {
	employee.Credit += credit
}

func (employee *Employee) RemoveCredits(credit float64) {
	employee.Credit -= credit
}

func (employee *Employee) CheckCredits() float64 {
	return employee.Credit
}

func (employee Employee) String() string {
	return employee.FirstName + employee.LastName
}
