package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	lastZ := 0.0
	for i := 0; math.Abs(z - lastZ) > 0.0000000001; i++ {
		lastZ = z
		z = z - (((z*z)-x)/(2*z))
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
