package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("negative value %v",float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x>=0.0 {return math.Sqrt(x),nil}
	var e ErrNegativeSqrt = ErrNegativeSqrt(x)
	return x, e
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
