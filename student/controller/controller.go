package controller

import (
	"log"
	"net/http"

	"github.com/MohitArora1/devops/student/student"
	"github.com/gorilla/mux"
)

//RunController for start server
func RunController(host string) {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/students", student.GetStudents).Methods("GET")
	r.HandleFunc("/api/v1/students", student.PostStudent).Methods("POST")
	log.Printf("starting server")
	err := http.ListenAndServe(host, r)
	if err != nil {
		log.Printf("error to start server :%v", err)
	}
}
