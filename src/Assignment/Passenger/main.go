package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// --- STRUCT ---
type Passenger struct {
	PassengerID int
	FirstName   string
	LastName    string
	MobileNo    int
	Email       string
}

var db *sql.DB

func InsertPassenger(db *sql.DB, passenger Passenger) {
	query := fmt.Sprintf("INSERT INTO Passengers VALUES ('%s', '%s', %d, %s)",
		passenger.FirstName, passenger.LastName, passenger.MobileNo, passenger.Email)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditPassenger(db *sql.DB, passenger Passenger) {
	query := fmt.Sprintf(
		"UPDATE Passengers SET FirstName='%s', LastName='%s', MobileNo=%d, Email='%s' WHERE PassengerID=%d",
		passenger.FirstName, passenger.LastName, passenger.MobileNo, passenger.Email, passenger.PassengerID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func DeletePassenger(db *sql.DB, PID int) {
	fmt.Println("Account cannot be deleted.")
}

func GetPassengerRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM assignment1.Passengers")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var passenger Passenger
		err = results.Scan(&passenger.PassengerID, &passenger.FirstName,
			&passenger.LastName, &passenger.MobileNo,
			&passenger.Email)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(passenger.PassengerID, passenger.FirstName,
			passenger.LastName, passenger.MobileNo,
			passenger.Email)
	}
}

func GetByPassengerID(db *sql.DB, passenger Passenger) (Passenger, bool) {
	query := fmt.Sprintf("SELECT * FROM assignment1.Passengers WHERE PassengerID=%d",
		passenger.PassengerID)

	results := db.QueryRow(query)

	var results2 = results.Scan(&passenger.PassengerID, &passenger.FirstName,
		&passenger.LastName, &passenger.MobileNo,
		&passenger.Email)

	var msg bool = false

	if results2 == sql.ErrNoRows {
		return passenger, msg
	}
	msg = true
	return passenger, msg
}

func CheckPassengerEmail(db *sql.DB, passenger Passenger) bool {
	query := fmt.Sprintf("SELECT Email FROM assignment1.Passengers")
	var matchEmail string

	results := db.QueryRow(query)

	results.Scan(&matchEmail)

	return matchEmail == passenger.Email

}

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

func UserPassenger(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}

	//params := mux.Vars(r)

	if r.Method == "GET" {
		var passenger Passenger
		var msg bool

		passenger, msg = GetByPassengerID(db, passenger)

		if msg {
			json.NewEncoder(w).Encode(passenger)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Passenger found"))
		}
	}

	if r.Method == "DELETE" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Account cannot be deleted"))
	}

	if r.Header.Get("Content-type") == "application/json" {
		// POST - Create new Passenger
		if r.Method == "POST" {
			var passenger Passenger
			var pid = strconv.Itoa(passenger.PassengerID)
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &passenger)

				if passenger.FirstName == "" || passenger.LastName == "" || passenger.Email == "" || passenger.MobileNo == 0 {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply information in JSON format"))
					return
				}

				var match = CheckPassengerEmail(db, passenger)

				if match {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Passenger ID"))
				} else {
					InsertPassenger(db, passenger)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Passenger added: " + pid))
				}

			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply passenger information in JSON format"))
			}
		}

		// PUT - Create or update existing Passenger
		if r.Method == "PUT" {
			var passenger Passenger
			var pid = strconv.Itoa(passenger.PassengerID)
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &passenger)

				if passenger.FirstName == "" || passenger.LastName == "" || passenger.Email == "" || passenger.MobileNo == 0 {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply information in JSON format"))
					return
				} else {
					EditPassenger(db, passenger)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte(
						"202 - Passenger updated: " + pid))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply passenger information in JSON format"))
			}
		}
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assignment1")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/passengers", UserPassenger).Methods("POST")
	router.HandleFunc("/api/v1/passengers/{passengerid}", UserPassenger).Methods(
		"GET", "PUT", "DELETE")

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
