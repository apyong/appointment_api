package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Appointment struc
type Appointment struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Place       string    `json:"place"`
	Participant string    `json:"participant"`
}

//Init Db when called
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbName := "appointment_db"
	db, err := sql.Open(dbDriver, dbUser+"@tcp(127.0.0.1:3306)/"+dbName+"?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	return db
}

//GetAppointments Sample
func GetAppointments(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	w.Header().Set("Content-Type", "application/json")

	app := Appointment{}
	finalResult := []Appointment{}

	stmt, err := db.Query("SELECT * FROM appointments")
	if err != nil {
		panic(err.Error())
	}

	//iterate on each appointments
	for stmt.Next() {

		err := stmt.Scan(&app.ID, &app.Name, &app.Description, &app.Date, &app.Place, &app.Participant)
		if err != nil {
			panic(err.Error())
		}

		finalResult = append(finalResult, app)
	}

	json.NewEncoder(w).Encode(finalResult)
	defer stmt.Close()

}

//GetAppointment Sample
func GetAppointment(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	params := mux.Vars(r)
	stmt, err := db.Query("SELECT * FROM appointments WHERE id=?", params["id"])

	if err != nil {
		panic(err.Error())
	}

	app := Appointment{}

	for stmt.Next() {
		var id int
		var name, description, place, participant string
		var date time.Time

		err = stmt.Scan(&id, &name, &description, &date, &place, &participant)

		if err != nil {
			panic(err.Error())
		}

		app.ID = id
		app.Name = name
		app.Description = description
		app.Date = date
		app.Place = place
		app.Participant = participant
	}

	json.NewEncoder(w).Encode(app)
	defer db.Close()

}

//CreateAppointment Sample
func CreateAppointment(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	w.Header().Set("Content-Type", "application/json")
	var app Appointment

	err := json.NewDecoder(r.Body).Decode(&app)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Prepare Sql Statement
	stmt, err := db.Prepare("INSERT INTO appointments(id,name, description, date, place, participant) VALUES ('',?,?,?,?,?)")

	if err != nil {
		panic(err.Error())
	}

	//Execute command
	res, err := stmt.Exec(app.Name, app.Description, app.Date, app.Place, app.Participant)
	lid, err := res.LastInsertId()

	app.ID = int(lid)

	//Show data if successfully inserted
	json.NewEncoder(w).Encode(app)

	defer db.Close()
}

//UpdateAppointment Sample
func UpdateAppointment(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	w.Header().Set("Content-Type", "application/json")
	var app Appointment
	params := mux.Vars(r)

	//Read body request
	err := json.NewDecoder(r.Body).Decode(&app)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Prepare Sql Statement
	stmt, err := db.Prepare("UPDATE appointments SET name=?, description=?, date=?, place=?, participant=? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(app.Name, app.Description, app.Date, app.Place, app.Participant, params["id"])

	i, err := strconv.Atoi(params["id"])
	app.ID = i

	//Show data if successfully updated
	json.NewEncoder(w).Encode(app)

	defer db.Close()
}

//DeleteAppointment Sample
func DeleteAppointment(w http.ResponseWriter, r *http.Request) {

	db := dbConn()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM appointments WHERE id=?")

	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(params["id"])

	response := map[string]string{
		"Message": "Successfuly Deleted!",
	}

	json.NewEncoder(w).Encode(response)
	defer db.Close()
}
