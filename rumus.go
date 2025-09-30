package main

import (
	"fmt"
	"math"
)

func main() {
	var x, z, a float64

	fmt.Scan(&x)
	a = x - 5
	z = 2/a - 5
	fmt.Println(int(math.Round(z)))

}
