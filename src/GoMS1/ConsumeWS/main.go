package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Success   bool
	timestamp int
	Base      string
	Date      string
	Rates     map[string]float64
}

var apis map[int]string

func fetchData(API int) {
	url := apis[API]
	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var result map[string]interface{}

			json.Unmarshal(body, &result)

			switch API {
			case 1:
				if result["success"] == true {
					fmt.Println(result["rates"].(map[string]interface{})["USD"])
				} else {
					fmt.Println(result["error"].(map[string]interface{})["info"])
				}

			case 2:
				if result["main"] != nil {
					fmt.Println(result["main"].(map[string]interface{})["temp"])
				} else {
					fmt.Println(result["message"])
				}
			}
		} else {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}

func main() {

	apis = make(map[int]string)

	apis[1] = "http://data.fixer.io/api/latest?access_key=740ea337f242cb7b84c7e0eefaf0519d"
	apis[2] =
		"http://api.openweathermap.org/data/2.5/weather?q=SINGAPORE&appid=7f7707cb6475370afdf155e98ed5d89f"

	go fetchData(1)
	go fetchData(2)

	fmt.Scanln()

	// url := "http://data.fixer.io/api/latest?access_key=740ea337f242cb7b84c7e0eefaf0519d"

	// if resp, err := http.Get(url); err == nil {
	// 	defer resp.Body.Close()
	// 	if body, err := ioutil.ReadAll(resp.Body); err == nil {
	// 		var result Result
	// 		json.Unmarshal(body, &result)
	// 		if result.Success {
	// 			// create a slice to store all keys
	// 			keys := make([]string, 0,
	// 				len(result.Rates))

	// 			// get all the keys---
	// 			for k := range result.Rates {
	// 				keys = append(keys, k)
	// 			}

	// 			// sort the keys
	// 			sort.Strings(keys)

	// 			// print the keys and values in
	// 			// alphabetical order
	// 			for _, k := range keys {
	// 				fmt.Println(k, result.Rates[k])
	// 			}
	// 			/*
	// 				for i, v := range result.Rates {
	// 					fmt.Println(i, v)
	// 				}
	// 			*/
	// 		} else {
	// 			var err Error
	// 			json.Unmarshal(body, &err)
	// 			fmt.Println(err.Error.Info)
	// 		}
	// 		//fmt.Println(string(body))
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Done")
}
