package main

import "C"
import (
	"fmt"
	"math"
)

func main() {
	// fmt.Println("Hello world!")
	var x0 = []float64{1., 2}
	fmt.Println(Derivative(testFunc, x0, 0.001))

}

//export doSome
func doSome() {
	println("Hello from Go")
}

//export getSomeNumber
func getSomeNumber() int8 {
	return 5
}

//export factorial
func factorial(n int64) int64 {
	var fact int64 = 1
	var i int64
	for i = 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

// export Derivative
// This function calculate derivative of function at point x0
func Derivative(function func(slice []float64) float64, x0 []float64, h float64) []float64 {
	if h == 0. {
		h = 1e-4
	}
	var xLen = len(x0)
	grad := x0

	for i := 0; i < xLen; i++ {
		xMinusH := x0
		xMinusH[i] -= h
		fMinus := function(xMinusH)
		xPlusH := x0
		xPlusH[i] += 2 * h
		fPlus := function(xPlusH)
		grad[i] = (fPlus - fMinus) / (2 * h)
		fmt.Println(grad, fPlus, fMinus)
	}
	return grad
}

func testFunc(x []float64) float64 {
	return math.Pow(x[0], 2) + math.Pow(x[1], 2)
}
