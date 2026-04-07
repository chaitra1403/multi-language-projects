package service

import "go-student-api/model"

var students []model.Student

func AddStudent(s model.Student) {
	students = append(students, s)
}

func GetAllStudents() []model.Student {
	return students
}

func GetStudentByID(id int) *model.Student {
	for i := range students {
		if students[i].ID == id {
			return &students[i]
		}
	}
	return nil
}

func DeleteStudent(id int) bool {
	for i := range students {
		if students[i].ID == id {
			students = append(students[:i], students[i+1:]...)
			return true
		}
	}
	return false
}
