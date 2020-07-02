package main

import (
	"log"
	"net/http"

	"github.com/apyong/appointment_api/httpd/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/appointments", handler.GetAppointments).Methods("GET")
	r.HandleFunc("/api/appointments/{id}", handler.GetAppointment).Methods("GET")
	r.HandleFunc("/api/appointments", handler.CreateAppointment).Methods("POST")
	r.HandleFunc("/api/appointments/{id}", handler.UpdateAppointment).Methods("PUT")
	r.HandleFunc("/api/appointments/{id}", handler.DeleteAppointment).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
