package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64, z float64, recursionDepth int, previousValue float64) (float64, error) { // wanted to make this recursive so argument list is a bit weird

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	if recursionDepth > 10 {
		return x, nil
	}

	newZ := z - (z*z-x)/(2*z)

	fmt.Printf("x %v z %v recursionDepth %v\n", x, newZ, recursionDepth)

	if math.Abs(previousValue-newZ) <= 0.1 {
		return newZ, nil // close enough
	} else {
		return Sqrt(x, newZ, recursionDepth+1, z)
	}
}

func main() {
	fmt.Println(Sqrt(2, 1, 0, 2))
	fmt.Println(Sqrt(-2, 1, 0, -2))
}
