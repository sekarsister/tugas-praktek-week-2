package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Masukan input Celsius: ")
	fmt.Scanln(&c)

	f := (c * 9 / 5) + 32
	k := c + 273.15
	r := c * 4 / 5

	fmt.Println("Suhu dalam Fahrenheit:", int(f))
	fmt.Println("Suhu dalam Kelvin:", int(k))
	fmt.Println("Suhu dalam Reamur:", int(r))
}
