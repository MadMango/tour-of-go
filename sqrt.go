package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, z float64, recursionDepth int, previousValue float64) float64 { // wanted to make this recursive so argument list is a bit weird
	if recursionDepth > 10 {
		return x
	}

	newZ := z - (z*z-x)/(2*z)

	fmt.Printf("x %v z %v recursionDepth %v\n", x, newZ, recursionDepth)

	if math.Abs(previousValue-newZ) <= 0.1 {
		return newZ // close enough
	} else {
		return Sqrt(x, newZ, recursionDepth+1, z)
	}
}

func main() {
	fmt.Printf("result %v", Sqrt(500, 1, 0, 500))
}
