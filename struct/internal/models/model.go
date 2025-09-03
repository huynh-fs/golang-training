package models

type Classes struct {
	Name          	string
	NumOfStudents 	int
}

type Students struct {
	Name      string
	Class     *Classes
}