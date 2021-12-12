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

type Driver struct {
	DriverID         int
	FirstName        string
	LastName         string
	MobileNo         int
	Email            string
	IdentificationNo string
	CarLicenseNo     string
}

var db *sql.DB

func InsertDriver(db *sql.DB, driver Driver) {
	query := fmt.Sprintf("INSERT INTO Drivers VALUES ('%s', '%s', %d, '%s', '%s', '%s')",
		driver.FirstName, driver.LastName, driver.MobileNo, driver.Email, driver.IdentificationNo, driver.CarLicenseNo)

	_, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}
}

func EditDriver(db *sql.DB, driver Driver) {
	query := fmt.Sprintf(
		"UPDATE Drivers SET FirstName='%s', LastName='%s', MobileNo=%d, Email='%s', IdentificationNo='%s', CarLicenseNo='%s' WHERE DriverID=%d",
		driver.FirstName, driver.LastName, driver.MobileNo, driver.Email, driver.IdentificationNo, driver.CarLicenseNo, driver.DriverID)
	_, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteDriver(db *sql.DB, DID int) {
	fmt.Println("Account cannot be deleted.")
}

func GetDriverRecords(db *sql.DB) {
	results, err := db.Query("Select * FROM assignment1_driver.Drivers")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		// map this type to the record in the table
		var driver Driver
		err = results.Scan(&driver.DriverID, &driver.FirstName,
			&driver.LastName, &driver.MobileNo,
			&driver.Email, &driver.IdentificationNo,
			&driver.CarLicenseNo)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(driver.DriverID, driver.FirstName,
			driver.LastName, driver.MobileNo,
			driver.Email, driver.IdentificationNo,
			driver.CarLicenseNo)
	}
}

func GetByDriverID(db *sql.DB, driver Driver) (Driver, bool) {
	query := fmt.Sprintf("SELECT * FROM assignment1_driver.Drivers WHERE DriverID=%d",
		driver.DriverID)

	results := db.QueryRow(query)

	var results2 = results.Scan(&driver.DriverID, &driver.FirstName,
		&driver.LastName, &driver.MobileNo,
		&driver.Email, &driver.IdentificationNo,
		&driver.CarLicenseNo)

	var msg bool = false

	if results2 == sql.ErrNoRows {
		return driver, msg
	}
	msg = true
	return driver, msg
}

func CheckDriverEmail(db *sql.DB, driver Driver) bool {
	query := fmt.Sprintf("SELECT Email FROM assignment1_driver.Drivers")
	var matchEmail string

	results := db.QueryRow(query)

	results.Scan(&matchEmail)

	return matchEmail == driver.Email

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

func UserDriver(w http.ResponseWriter, r *http.Request) {
	if !validKey(r) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("401 - Invalid key"))
		return
	}

	params := mux.Vars(r)

	if r.Method == "GET" {
		var driver Driver
		var msg bool

		driver, msg = GetByDriverID(db, driver)

		if msg {
			json.NewEncoder(w).Encode(driver)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - No Driver found"))
		}
	}

	if r.Method == "DELETE" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Account cannot be deleted"))
	}

	if r.Header.Get("Content-type") == "application/json" {
		// POST - Create new Driver
		if r.Method == "POST" {
			var driver Driver
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &driver)

				if driver.FirstName == "" || driver.LastName == "" || driver.Email == "" || driver.MobileNo == 0 || driver.IdentificationNo == "" || driver.CarLicenseNo == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply information in JSON format"))
					return
				}

				var match = CheckDriverEmail(db, driver)

				if match {
					w.WriteHeader(http.StatusConflict)
					w.Write([]byte(
						"409 - Duplicate Driver ID"))
				} else {
					InsertDriver(db, driver)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte("201 - Driver added: " + params["driverid"]))
				}

			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply driver information in JSON format"))
			}
		}

		// PUT - Create or update existing Driver
		if r.Method == "PUT" {
			var driver Driver
			reqBody, err := ioutil.ReadAll(r.Body)

			if err == nil {
				json.Unmarshal(reqBody, &driver)

				if driver.FirstName == "" || driver.LastName == "" || driver.Email == "" || driver.MobileNo == 0 || driver.IdentificationNo == "" || driver.CarLicenseNo == "" {
					w.WriteHeader(http.StatusUnprocessableEntity)
					w.Write([]byte(
						"422 - Please supply information in JSON format"))
					return
				} else {
					EditDriver(db, driver)
					w.WriteHeader(http.StatusAccepted)
					w.Write([]byte(
						"202 - Driver updated: " + params["driverid"]))
				}
			} else {
				w.WriteHeader(
					http.StatusUnprocessableEntity)
				w.Write([]byte("422 - Please supply driver information in JSON format"))
			}
		}
	}
}

func main() {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/assignment1_driver")

	// handle error
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/api/drivers", UserDriver).Methods("POST")
	router.HandleFunc("/api/drivers/{driverid}", UserDriver).Methods(
		"GET", "PUT", "DELETE")

	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
