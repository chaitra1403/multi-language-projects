package handler

import (
	"encoding/json"
	"go-student-api/model"
	"go-student-api/service"
	"net/http"
	"strconv"
)

func AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	var s model.Student
	json.NewDecoder(r.Body).Decode(&s)

	service.AddStudent(s)
	json.NewEncoder(w).Encode(s)
}

func GetAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(service.GetAllStudents())
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	student := service.GetStudentByID(id)
	if student == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(student)
}

func DeleteStudentHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idParam)

	if service.DeleteStudent(id) {
		w.Write([]byte("Deleted"))
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
