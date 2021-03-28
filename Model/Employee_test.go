package Model

import (
	"testing"
)

var employee Employee

func InitEmployee() {
	employee = Employee{
		Account: Account{
			FirstName: "kazuya",
			LastName:  "sugimoto",
		},
		Credit: 100,
	}
}

func Test_Employee_ChangeName(t *testing.T) {
	InitEmployee()
	employee.ChangeName("update", "name")

	expect := "updatename"
	actual := employee.Account.FirstName + employee.Account.LastName

	if expect != actual {
		t.Error("Does not much ", expect, actual)
	}
}

func Test_Employee_AddCredtis(t *testing.T) {
	InitEmployee()

	employee.AddCredits(100)

	var expect float64
	expect = 200
	actual := employee.Credit

	if expect != actual {
		t.Error("Does not much ", expect, actual)
	}
}

func Test_Employee_RemoveCredits(t *testing.T) {
	InitEmployee()

	employee.RemoveCredits(50)

	var expect float64
	expect = 50
	actual := employee.Credit

	if expect != actual {
		t.Error("Does not much ", expect, actual)
	}
}

func Test_Employee_CheckCredit(t *testing.T) {
	InitEmployee()

	var expect float64
	expect = 100
	actual := employee.CheckCredits()

	if expect != actual {
		t.Error("Does not much ", expect, actual)
	}
}

func Test_Employee_String(t *testing.T) {
	InitEmployee()

	expect := "kazuyasugimoto"
	actual := employee.String()

	if expect != actual {
		t.Error("Does not much ", expect, actual)
	}

}
