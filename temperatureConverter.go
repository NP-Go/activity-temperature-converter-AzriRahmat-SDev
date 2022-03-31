package main

import "fmt"

var resultFahrenheit float64
var resultCelsius float64
var resultKelvin float64

func convertKelvin(temperatureInput float64) (float64, float64) {
	resultFahrenheit = (temperatureInput * (9.0 / 5.0)) - 459.67
	resultCelsius = (temperatureInput - 32) * (5.0 / 9.0)
	return resultFahrenheit, resultCelsius
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	resultFahrenheit = temperatureInput*(9.0/5.0) + 32
	resultKelvin = (5.0 / 9.0) * (resultFahrenheit + 459.67)
	return resultFahrenheit, resultKelvin
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	resultKelvin = (5.0 / 9.0) * (resultFahrenheit + 459.67)
	resultCelsius = (temperatureInput - 32) * (5.0 / 9.0)
	return resultKelvin, resultCelsius
}

func main() {
	var temperatureChoice int
	var temperatureInput float64
	fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
	fmt.Scanln(&temperatureChoice)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)
	if temperatureChoice == 1 {
		convertKelvin(temperatureInput)
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Celsius: ", resultCelsius)
		fmt.Printf("Fahrenheit: %0.2f\n", resultFahrenheit)
		fmt.Printf("Celsius: %0.2f\n", resultCelsius)
	}
	if temperatureChoice == 2 {
		convertCelsius(temperatureInput)
		fmt.Println("Fahrenheit: ", resultFahrenheit, " Kelvin: ", resultKelvin)
	}
	if temperatureChoice == 3 {
		convertFahrenheit(temperatureInput)
		fmt.Println("Kelvin: ", resultKelvin, " Celsius: ", resultCelsius)
	}

}
