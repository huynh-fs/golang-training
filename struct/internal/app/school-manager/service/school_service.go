package service

import (
	"fmt"
	"github.com/huynh-fs/struct/internal/domain"
)

type SchoolService struct {
	classes  []domain.Classes
	students []domain.Students
}

func NewSchoolService() *SchoolService {
	return &SchoolService{
		classes:  make([]domain.Classes, 0),
		students: make([]domain.Students, 0),
	}
}

func (s *SchoolService) CreateClass(name string) error {
	for _, class := range s.classes {
		if class.Name == name {
			return fmt.Errorf("lớp '%s' đã tồn tại", name)
		}
	}
	newClass := domain.Classes{Name: name, NumOfStudents: 0}
	s.classes = append(s.classes, newClass)
	return nil
}


func (s *SchoolService) AddStudent(studentName, className string) error {
	classIndex := -1
	for i := range s.classes {
		if s.classes[i].Name == className {
			classIndex = i
			break
		}
	}

	if classIndex == -1 {
		return fmt.Errorf("lớp '%s' không tồn tại", className)
	}

	newStudent := domain.Students{Name: studentName, Class: className}
	s.students = append(s.students, newStudent)

	s.classes[classIndex].NumOfStudents++

	return nil
}

func (s *SchoolService) GetAllClasses() []domain.Classes {
	return s.classes
}

func (s *SchoolService) GetStudentsByClassName(className string) []domain.Students {
	var studentsInClass []domain.Students
	for _, student := range s.students {
		if student.Class == className {
			studentsInClass = append(studentsInClass, student)
		}
	}
	return studentsInClass
}