package service

import (
	"fmt"
	"github.com/huynh-fs/struct/internal/domain"
)

type ClassService struct {
	classes []domain.Classes
}

func NewClassService() *ClassService {
	return &ClassService{
		classes: make([]domain.Classes, 0),
	}
}

func (s *ClassService) CreateClass(name string) error {
	for _, class := range s.classes {
		if class.Name == name {
			return fmt.Errorf("lớp '%s' đã tồn tại", name)
		}
	}
	newClass := domain.Classes{
		Name:     name,
		Students: make([]*domain.Students, 0),
	}
	s.classes = append(s.classes, newClass)
	return nil
}

func (s *ClassService) AddStudentToClass(studentName, className string) error {
	var targetClass *domain.Classes
	for i := range s.classes {
		if s.classes[i].Name == className {
			targetClass = &s.classes[i]
			break
		}
	}
	if targetClass == nil {
		return fmt.Errorf("lớp '%s' không tồn tại", className)
	}

	student := &domain.Students{Name: studentName, Class: targetClass}
	targetClass.AddStudent(student)
	return nil
}

func (s *ClassService) GetAllClasses() []domain.Classes {
	return s.classes
}