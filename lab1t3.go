package main

import (
	"fmt"
	"strings"
)

type Student struct {
	rollnumber int
	name       string
	address    string
}

func NewStudent(rollno int, name string, address string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {
	st := NewStudent(rollno, name, address)

	ls.list = append(ls.list, st)

	return st
}
func display(stlist *StudentList) {
	for i := 0; i < len(stlist.list); i++ {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Printf("RollNo: %d\n", stlist.list[i].rollnumber)
		fmt.Printf("Name: %s\n", stlist.list[i].name)
		fmt.Printf("Address: %s\n", stlist.list[i].address)
	}
}
func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBBBBB")
	student.CreateStudent(26, "Aisha", "CCCCCCCCC")
	display(student)
}
