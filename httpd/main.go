package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/apyong/appointment_api/httpd/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	// createDB()

	r := mux.NewRouter()
	r.HandleFunc("/api/appointments", handler.GetAppointments).Methods("GET")
	r.HandleFunc("/api/appointments/{id}", handler.GetAppointment).Methods("GET")
	r.HandleFunc("/api/appointments", handler.CreateAppointment).Methods("POST")
	r.HandleFunc("/api/appointments/{id}", handler.UpdateAppointment).Methods("PUT")
	r.HandleFunc("/api/appointments/{id}", handler.DeleteAppointment).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func createDB() {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}

	_, err = db.Exec("CREATE DATABASE appointment_db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}

	_, err = db.Exec("USE appointment_db")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	stmt, err := db.Prepare("CREATE TABLE appointments(id int NOT NULL AUTO_INCREMENT, name varchar(120), description varchar(255), date datetime, place varchar(120), participant varchar(120), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully.. ")
	}

	defer db.Close()
}
