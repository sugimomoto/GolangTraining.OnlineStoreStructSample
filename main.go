package main

import (
	. "OnlineStoreStructSample/Model"
	"fmt"
)

func main() {
	employee := Employee{
		Account: Account{
			FirstName: "kazuya",
			LastName:  "sugimoto",
		},
		Credit: 100,
	}

	employee.AddCredits(100)
	employee.RemoveCredits(50)
	employee.ChangeName("update", "name")

	fmt.Println(employee, employee.CheckCredits())
}
