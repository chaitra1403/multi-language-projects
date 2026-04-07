package main

import (
	"fmt"
	"go-student-api/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/add", handler.AddStudentHandler)
	http.HandleFunc("/students", handler.GetAllStudentsHandler)
	http.HandleFunc("/student", handler.GetStudentHandler)
	http.HandleFunc("/delete", handler.DeleteStudentHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
