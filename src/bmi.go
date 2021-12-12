package main

import "fmt"

func main() {
	var weight, height, bmi float32

	fmt.Print("Enter Weight (Kg): ")
	n, err := fmt.Scanf("%f\n", &weight)
	if err != nil || n != 1 {
		fmt.Println(n, err)
	}

	fmt.Print("Enter Height (m): ")
	n, err = fmt.Scanf("%f\n", &height)
	if err != nil || n != 1 {
		fmt.Println(n, err)
	}

	bmi = weight / (height * height)
	fmt.Println("BMI: ", bmi)
}
