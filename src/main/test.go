package main

/*
types
	int (5, -200)
	int64
	float32 (1.7, -2.6)
	float64

slices: arrays of another type
maps: key vaule pairs of 2 types

functions:
func main, factorial (...) (int, ...) {
	var jdkal int
	jdkal = 5

	foo := 5
	
	x, y, z, := 3, "hello", 5.3	//fine
	w, z := 10, 20			//also fine
	......

}

variables

constants
*/

import (
	"fmt"
)

func factorial(base int) int {
	//5! 120
	//3! 6
	//1! 1
	//0! 0

	//3! 3 * 2 * 1
	//5! 5*4*3*2*1

	var output int = 1
	for i := 1; i <= base; i ++ {
		output *= i 
	}
	return output
}

func main() {
	fmt.Println(factorial(5))
}

