package main

import "fmt"

func main() {
	var (
		f int
	)
	fmt.Print("Tahun ")
	fmt.Scanln(&f)
	if f%4 == 0 && (f%100 != 0 || f%400 == 0) {
		fmt.Printf("%d adalah tahun kabisat\n", f)
	} else {
		fmt.Printf("%d bukan tahun kabisat\n", f)
	}
}
