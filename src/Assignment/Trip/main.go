package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Trip struct {
	TripID      int
	Pickup      string
	Dropoff     string
	Assigned    string
	PassengerID int
	DriverID    int
}

var db *sql.DB

func InsertTrip(db *sql.DB, trip Trip) {
	query := fmt.Sprintf("INSERT INTO Trips VALUES ('%s', '%s', %d, %d)",
		trip.Pickup, trip.Dropoff, trip.PassengerID, trip.DriverID)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditTrip(db *sql.DB, trip Trip) {
	query := fmt.Sprintf(
		"UPDATE Trips SET Pickup='%s', Dropoff='%s', Assigned='%s', PassengerID=%d, driverID=%d WHERE TripID=%d",
		trip.Pickup, trip.Dropoff, trip.Assigned, trip.PassengerID, trip.DriverID, trip.TripID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteTrip(db *sql.DB, TID int) {
	fmt.Println("Trip cannot be deleted.")
}

func GetPassengerTripRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM assignment1_trip.Trips WHERE ")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var trip Trip
		err = results.Scan(&trip.TripID, &trip.Pickup,
			&trip.Dropoff, &trip.Assigned,
			&trip.PassengerID, &trip.DriverID)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(trip.TripID, trip.Pickup,
			trip.Dropoff, trip.Assigned,
			trip.PassengerID, trip.DriverID)
	}
}

func GetByTripID(db *sql.DB, trip Trip) (Trip, bool) {
	query := fmt.Sprintf("SELECT * FROM assignment1_trip.Trips WHERE TripID=%d",
		trip.TripID)

	results := db.QueryRow(query)

	var results2 = results.Scan(&trip.TripID, &trip.Pickup,
		&trip.Dropoff, &trip.Assigned,
		&trip.PassengerID, &trip.DriverID)

	var msg bool = false

	if results2 == sql.ErrNoRows {
		return trip, msg
	}
	msg = true
	return trip, msg
}

// func CheckTripEmail(db *sql.DB, driver Driver) bool {
// 	query := fmt.Sprintf("SELECT Email FROM assignment1_driver.Drivers")
// 	var matchEmail string

// 	results := db.QueryRow(query)

// 	results.Scan(&matchEmail)

// 	return matchEmail == driver.Email

// }
func validKey(r *http.Request) bool {
	v := r.URL.Query()
	if key, ok := v["key"]; ok {
		if key[0] == "2c78afaf-97da-4816-bbee-9ad239abb296" {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func PassengerTrip(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}

	// params := mux.Vars(r)

	if r.Method == "GET" {
		var trip Trip
		var msg bool

		trip, msg = GetByTripID(db, trip)

		if msg {
			json.NewEncoder(w).Encode(trip)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Trip found"))
		}
	}

	if r.Method == "DELETE" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Trip cannot be deleted"))
	}

	if r.Header.Get("Content-type") == "application/json" {
		// POST - Create new Trip
		if r.Method == "POST" {
			var trip Trip
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &trip)

				if trip.Pickup == "" || trip.Dropoff == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply information in JSON format"))
					return
				}

				// var match = CheckTripEmail(db, driver)

				// if match {
				// 	w.WriteHeader(http.StatusConflict)
				// 	w.Write([]byte(
				// 		"409 - Duplicate Trip ID"))
				// } else {
				// 	InsertTrip(db, trip)
				// 	w.WriteHeader(http.StatusCreated)
				// 	w.Write([]byte("201 - Trip added: " + strconv.Itoa(trip.TripID)))
				// }

			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply driver information in JSON format"))
			}
		}

		// PUT - Create or update existing Driver
		if r.Method == "PUT" {

		}
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assignment1_trip")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api/trips", GetPassengerTripRecords).Methods("GET")
	router.HandleFunc("/api/trips", PassengerTrip).Methods("POST")
	router.HandleFunc("/api/trips/{tripid}", PassengerTrip).Methods(
		"GET", "PUT", "DELETE")

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
