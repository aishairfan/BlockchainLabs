package main

import (
	"crypto/sha256"
	"fmt"
)

func CalculateHash(strtoHash string) string {
	fmt.Printf("String Received:     %s\n", strtoHash)
	fmt.Print("HASH:     ")
	return fmt.Sprintf("%x", sha256.Sum256([]byte(strtoHash)))
}

type Student struct {
	rollnumber int
	name       string
	address    string
	subjects   []string
}

func NewStudent(rollno int, name string, address string, sbjlist []string) *Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	for i := 0; i < len(sbjlist); i++ {
		s.subjects = append(s.subjects, sbjlist[i])
	}
	return s
}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, sbjlist []string) *Student {
	st := NewStudent(rollno, name, address, sbjlist)

	ls.list = append(ls.list, st)
	return st
}
func getData(stdnt *Student) string {
	var final string
	final = fmt.Sprintf("%s %d %s", stdnt.name, stdnt.rollnumber, stdnt.address)
	for i := 0; i < len(stdnt.subjects); i++ {
		final = final + " " + stdnt.subjects[i]
	}
	return final

}
func main() {
	sbj := []string{"Phy", "Chem", "Bio"}
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAAA", sbj)
	student.CreateStudent(25, "Naveed", "BBBBBBBBB", sbj)
	student.CreateStudent(26, "Aisha", "CCCCCCCCC", sbj)
	for i := 0; i < len(student.list); i++ {
		fmt.Println(" ")
		fmt.Println(CalculateHash(getData(student.list[i])))
	}

}
