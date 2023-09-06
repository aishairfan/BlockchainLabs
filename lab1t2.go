package main

import "fmt"

type employee struct {
	name     string
	position string
	salary   int
}
type company struct {
	companyName string
	employees   []employee
}

func main() {
	emp1 := employee{"Amir", "Front-End Developer", 80000}
	emp2 := employee{"Ahmed", "Back-End Developer", 90000}
	emp3 := employee{"Amna", "Business Analyst", 80000}
	emplys := []employee{emp1, emp2, emp3}

	var tetra company
	tetra.companyName = "Tetra"
	tetra.employees = emplys

	//printing company details
	fmt.Printf("Company Name: %s\n", tetra.companyName)
	for i := 0; i < len(tetra.employees); i++ {
		fmt.Printf("Employee=%d \nName=%s \nPosition: %s \nSalary=%d\n\n", i+1, tetra.employees[i].name, tetra.employees[i].position, tetra.employees[i].salary)
	}
}
