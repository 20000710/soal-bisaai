package main

import (
	"fmt"
	"math"
)

// soal a menghitung luas persegi
func calculateSquare(num int) int {
	return num * num
}

// soal b menghitung luas segitiga
func calculateTriangle(base, height int) int {
	return (base * height) / 2
}

// soal c menghitung luas lingkaran
func calculateCircle(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

// soal d menghitung sudut sinus
func calculateSin(angle float64) float64 {
	return math.Sin(angle)
}

// soal d menghitung sudut cosinus
func calculateCos(angle float64) float64 {
	return math.Cos(angle)
}

// soal d menghitung sudut tangen
func calculateTan(angle float64) float64 {
	return math.Tan(angle)
}

// soal e fungsi untuk menghitung jarak tempuh dalam gerak lurus beraturan
func calculateUniformMotion(speed, time float64) float64 {
	return speed * time
}

// soal f fungsi untuk menghitung jarak tempuh dalam gerak lurus berubah beraturan
func calculateUniformAccelerationMotion(initialVelocity, acceleration, time float64) float64 {
	return initialVelocity*time + 0.5*acceleration*math.Pow(time, 2)
}

// soal g fungsi untuk menghitung energi potensial
func calculatePotentialEnergy(mass, height float64) float64 {
	return mass * 9.8 * height
}

// soal g fungsi untuk menghitung energi kinetik
func calculateKineticEnergy(mass, velocity float64) float64 {
	return 0.5 * mass * math.Pow(velocity, 2)
}

// soal h fungsi untuk menghitung frekuensi atau getaran
func calculateFrequency(period float64) float64 {
	return 1 / period
}

// soal i fungsi untuk menghitung masa jenis
func calculateDensity(mass, volume float64) float64 {
	return mass / volume
}

// soal j fungsi untuk menghitung daya
func calculatePower(work, time float64) float64 {
	return work / time
}

// soal j fungsi untuk menghitung tekanan
func calculatePressure(force, area float64) float64 {
	return force / area
}

// soal j fungsi untuk menghitung usaha
func calculateEffort(force, distance float64) float64 {
	return force * distance
}

// soal j fungsi untuk menghitung gaya
func calculateForce(mass, acceleration float64) float64 {
	return mass * acceleration
}

// soal k fungsi untuk menghitung suhu dalam satuan Celcius ke Fahrenheit
func calculateCelciusToFahrenheit(celcius float64) float64 {
	return (celcius * 9 / 5) + 32
}

// soal k fungsi untuk menghitung suhu dalam satuan Fahrenheit ke Celcius
func calculateFahrenheitToCelcius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

// soal k fungsi untuk menghitung suhu dalam satuan Celcius ke Kelvin
func calculateCelciusToKelvin(celcius float64) float64 {
	return celcius + 273.15
}

// soal k fungsi untuk menghitung suhu dalam satuan Kelvin ke Celcius
func calculateKelvinToCelcius(kelvin float64) float64 {
	return kelvin - 273.15
}

// soal k fungsi untuk menghitung suhu dalam satuan Fahrenheit ke Kelvin
func calculateFahrenheitToKelvin(fahrenheit float64) float64 {
	celcius := calculateFahrenheitToCelcius(fahrenheit)
	return calculateCelciusToKelvin(celcius)
}

// soal k fungsi untuk menghitung suhu dalam satuan Kelvin ke Fahrenheit
func calculateKelvinToFahrenheit(kelvin float64) float64 {
	celcius := calculateKelvinToCelcius(kelvin)
	return calculateCelciusToFahrenheit(celcius)
}

func main() {
	result := calculateSquare(5)
	fmt.Println(result)

	triangleArea := calculateTriangle(4, 3)
	fmt.Println(triangleArea)
}


