package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {

	var result float64

	for z := 1.0; z < 10; z++ {
		if z == (z-((math.Pow(z, 2)) -x) / (2*z) ) {
			result = z
		} else {
			continue 
		}
		
	}

	return result
}

func main() {
    fmt.Println(Sqrt(49))


}