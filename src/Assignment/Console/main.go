package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Passenger struct {
	PassengerID int
	FirstName   string
	LastName    string
	MobileNo    int
	Email       string
}

type Driver struct {
	DriverID         int
	FirstName        string
	LastName         string
	MobileNo         int
	Email            string
	IdentificationNo string
	CarLicenseNo     string
}

type Trip struct {
	TripID      int
	Pickup      string
	Dropoff     string
	Assigned    string
	PassengerID int
	DriverID    int
}

const driversBaseURL = "http://localhost:5000/api/drivers"
const passengersBaseURL = "http://localhost:5000/api/passengers"
const tripsBaseURL = "http://localhost:5000/api/trips"
const key = "2c78afaf-97da-4816-bbee-9ad239abb296"

var userinput = ""

func Menu() {
	fmt.Println("---Main Menu---")
	fmt.Println("1. Create Passenger Account")
	fmt.Println("2. Create Driver Account")
	fmt.Println("3. Passenger Login")
	fmt.Println("4. Driver Login")
	fmt.Println("0. Exit Menu")
	fmt.Scanln(&userinput)

	switch userinput {
	case "1":
		CreatePassengerAccount()
	case "2":
		CreateDriverAccount()
	case "3":
		LoginPassenger()
	case "4":
		LoginDriver()
	case "0":
		break
	default:
		fmt.Println("Please enter a valid number")
		Menu()
	}
}

func LoginPassenger() {
	fmt.Println("---Passenger Login---")
	fmt.Println("Enter your Passenger ID:")
	fmt.Scanln(&userinput)

	//if success, go to passenger menu
	PassengerMenu()
}

func LoginDriver() {
	fmt.Println("---Driver Login---")
	fmt.Println("Enter your Driver ID:")
	fmt.Scanln(&userinput)

	//if success, go to driver menu
	DriverMenu()
}

func CreatePassengerAccount() {
	var firstName, lastName, mobileNo, email string
	id := 1

	fmt.Println("---Create Passenger Account---")
	fmt.Println("Enter First Name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter Mobile No:")
	fmt.Scanln(&mobileNo)
	fmt.Println("Enter Email:")
	fmt.Scanln(&email)

	jsonData := map[string]string{"FirstName": "", "LastName": "", "MobileNo": "", "Email": ""}
	jsonData["FirstName"] = firstName
	jsonData["LastName"] = lastName
	jsonData["MobileNo"] = mobileNo
	jsonData["Email"] = email

	addPassenger(strconv.Itoa(id), jsonData)
	id += 1
	//if success, go to passenger menu
	PassengerMenu()
}

func CreateDriverAccount() {
	id := 1
	var firstName, lastName, mobileNo, email, identificationNo, carLicenseNo string

	fmt.Println("---Create Driver Account---")
	fmt.Println("Enter First Name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter Mobile No:")
	fmt.Scanln(&mobileNo)
	fmt.Println("Enter Email:")
	fmt.Scanln(&email)
	fmt.Println("Enter Identification No:")
	fmt.Scanln(&identificationNo)
	fmt.Println("Enter Car License No:")
	fmt.Scanln(&carLicenseNo)

	jsonData := map[string]string{"FirstName": "", "LastName": "", "MobileNo": "", "EmailAddress": "", "IdentificationNo": "", "CarLicenseNo": ""}
	jsonData["FirstName"] = firstName
	jsonData["LastName"] = lastName
	jsonData["MobileNo"] = mobileNo
	jsonData["Email"] = email
	jsonData["IdentificationNo"] = identificationNo
	jsonData["CarLicenseNo"] = carLicenseNo

	addDriver(strconv.Itoa(id), jsonData)
	id += 1
	//if success, go to driver menu
	DriverMenu()

}

func PassengerMenu() {
	fmt.Println("---Passenger Menu---")
	fmt.Println("Passenger ID:" + "1") //passenger id
	fmt.Println("1. Update Information")
	fmt.Println("2. Request a Trip")
	fmt.Println("3. View All Trips")
	fmt.Println("0. Logout")
	fmt.Scanln(&userinput)

	switch userinput {
	case "1":
		UpdatePassengerInformation()
	case "2":
		RequestATrip()
	case "3":
		ViewAllTrips()
	case "0":
		Menu()
	default:
		fmt.Println("Please enter a valid number")
		Menu()
	}

}

func UpdatePassengerInformation() {
	var firstName, lastName, mobileNo, email string
	id := 1

	fmt.Println("Enter First Name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter Mobile No:")
	fmt.Scanln(&mobileNo)
	fmt.Println("Enter Email:")
	fmt.Scanln(&email)

	jsonData := map[string]string{"FirstName": "", "LastName": "", "MobileNo": "", "Email": ""}
	jsonData["FirstName"] = firstName
	jsonData["LastName"] = lastName
	jsonData["MobileNo"] = mobileNo
	jsonData["Email"] = email

	updatePassenger(strconv.Itoa(id), jsonData)
	id += 1

	PassengerMenu()

}

func RequestATrip() {

	fmt.Println("---Request a Trip---")
	fmt.Println("Pickup Postal Code:")
	fmt.Scanln(&userinput)
	fmt.Println("Dropoff Postal Code:")
	fmt.Scanln(&userinput)

	fmt.Println("Trip Requested! Going back to Passenger Menu...")
	PassengerMenu()
}

func ViewAllTrips() {
	fmt.Println("---View All Trips---")
	getTrips("")
}

func DriverMenu() {
	fmt.Println("---Driver Menu---")
	fmt.Println("Driver ID:" + "!") //driver id
	fmt.Println("1. Update Information")
	fmt.Println("2. Start Trip")
	fmt.Println("0. Logout")
	fmt.Scanln(&userinput)

	switch userinput {
	case "1":
		UpdateDriverInformation()
	case "2":
		DriverTrip()
	case "0":
		Menu()
	default:
		fmt.Println("Please enter a valid number")
		DriverMenu()
	}
}

func UpdateDriverInformation() {
	fmt.Println("---Update Driver Information---")
	// UpdateInformation()
	// fmt.Println("5. Identification No")
	// fmt.Println("6. Car License No")
	// fmt.Println("0. Back to Driver Menu")
	// fmt.Scanln(&userinput)

	id := 1
	var firstName, lastName, mobileNo, email, identificationNo, carLicenseNo string

	fmt.Println("Enter First Name:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter Mobile No:")
	fmt.Scanln(&mobileNo)
	fmt.Println("Enter Email:")
	fmt.Scanln(&email)
	fmt.Println("Enter Identification No:")
	fmt.Scanln(&identificationNo)
	fmt.Println("Enter Car License No:")
	fmt.Scanln(&carLicenseNo)

	jsonData := map[string]string{"FirstName": "", "LastName": "", "MobileNo": "", "EmailAddress": "", "IdentificationNo": "", "CarLicenseNo": ""}
	jsonData["FirstName"] = firstName
	jsonData["LastName"] = lastName
	jsonData["MobileNo"] = mobileNo
	jsonData["Email"] = email
	jsonData["IdentificationNo"] = identificationNo
	jsonData["CarLicenseNo"] = carLicenseNo

	updateDriver(strconv.Itoa(id), jsonData)
	id += 1
}

func DriverTrip() {
	//display all unassigned trips
	fmt.Println("---Driver Trip Menu---")
	fmt.Println("Trip ID:" + "1") //trip id
	fmt.Println("0. End Trip")

	if userinput == "0" {
		DriverMenu()
	}

	fmt.Scanln(&userinput)
}

func UpdateInformation() {
	fmt.Println("1. First Name")
	fmt.Println("2. Last Name")
	fmt.Println("3. Mobile No")
	fmt.Println("4. Email")
}

func getTrips(code string) {
	url := passengersBaseURL
	if code != "" {
		url = tripsBaseURL + "/" + code + "?key=" + key
	}
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func addPassenger(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(passengersBaseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func addDriver(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(driversBaseURL+"/"+code+"?key="+key,
		"application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func updatePassenger(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPut,
		passengersBaseURL+"/"+code+"?key="+key,
		bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}
func updateDriver(code string, jsonData map[string]string) {
	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest(http.MethodPut,
		driversBaseURL+"/"+code+"?key="+key,
		bytes.NewBuffer(jsonValue))

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(response.StatusCode)
		fmt.Println(string(data))
		response.Body.Close()
	}
}

func main() {
	Menu()
}
