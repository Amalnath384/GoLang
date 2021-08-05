package router

import (
	. "awesomeProject/Practice28-Mockery-Rest-Mongo/pkg/handler/student"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {

	studentHandler := NewStudentService()
	r := mux.NewRouter()

	r.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	r.HandleFunc("/students", studentHandler.ListStudent).Methods("GET")
	r.HandleFunc("/students/{id}", studentHandler.GetStudent).Methods("GET")
	r.HandleFunc("/students", studentHandler.CreateStudent).Methods("POST")
	r.HandleFunc("/students/{id}", studentHandler.UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", studentHandler.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/login", studentHandler.Login).Methods("POST")
	r.HandleFunc("/home", studentHandler.Home).Methods("GET")
	r.HandleFunc("/refresh", studentHandler.Refresh).Methods("GET")

	http.ListenAndServe("127.0.0.1:8080", r)

	return r
}
