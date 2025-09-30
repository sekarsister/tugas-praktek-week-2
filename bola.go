package main

import "fmt"

func main() {
	var (
		f float64
	)
	fmt.Print("Jejari ")
	fmt.Scanln(&f)
	var (
		t  float64 = (f * f * f) * 4 / 3
		w  float64 = (f * f) * 4
		pi float64 = 3.1415926535
	)
	fmt.Print("bola memiliki Volume ", float64(t*pi), " dan luas kulit", float64(w*pi))

}
