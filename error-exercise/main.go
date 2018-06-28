package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// Note:
	// When printing an error, go will call e.Error() to
	// get the string
	//
	// If we do not convert e to a float below, we will get an infinite loop.
	// Sprintf calls e.Error() to get the string representation.. which calls
	// this function... which calls Sprintf...
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(16))
	// outputs:
	// 4 <nil>, float64 and error
	fmt.Println(Sqrt(-2))
	// outputs:
	// 0 cannot Sqrt negative number: -2
}
