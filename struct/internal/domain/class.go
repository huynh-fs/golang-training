package domain

import (
	"fmt"
)

type Classes struct {
	Name 			string
	Students 		[]*Students
}

func (c *Classes) AddStudent(s *Students) {
	c.Students = append(c.Students, s)
}

func (c *Classes) NumOfStudents() int {
	return len(c.Students)
}

func (c *Classes) Display() {
	fmt.Printf("\n---------Lớp: %s | Số lượng sinh viên: %d ----------\n", c.Name, c.NumOfStudents())
	if c.NumOfStudents() == 0 {
		fmt.Println("Lớp chưa có sinh viên nào.")
	} else {
		for i, student := range c.Students {
			fmt.Printf("\t%d. %s\n", i+1, student.Name)
		}
	}
}