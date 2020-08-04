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

type MyInt int // making alias types

type MyStruct struct {
	foo int
	bar string
}

// having this function means MyStruct implements the "MyInterface" iface
func (a MyStruct) Foo() bool {
	return a.bar == "cheese"
	// name.member accesses a struct's members
}

// general format of variable declaration:
// "var" [name] [type]

// general format of a function declaration:
// "func" [name]([params]) ([return types]) {
//     [body]
// }
// name is its chosen name, params are variable declarations
// (without the "var"), return types are a list of types, and
// body is a list of statements

// general format of a type declaration:
// "type [name] [underlying]
// name is just our name for the type
// underlying is the specification for a legal name (often a struct)

type MyInterface interface {
	Foo() bool // this looks like a function declaration!
	// "func" is implied, and the body is left out
}

func main() {
	fmt.Println(factorial(5))

//	var a int
	// "named" and "underlying" type, both int

//	var x MyInt
	// what type is x
	// the "named type: MyInt
	// the "underlying" type: int

//	y := x + 5
	// I can use any type as though it were its UNDERLYING type
	// in terms of expressions and stuff

//	var z MyStruct
	// named type: MyStruct
	// underlying type: struct {
	// 	foo int
	// 	bar string
	//}
	// interfaces it implements, which are determined by 
	// functions it has.

	// interfaces
	// interfaces are lists of methods. all types implement at
	// least the "empty interface', interface{}

	var foo, foo2 MyInterface
	foo = MyStruct{} // creates a new MyStruct with default vaules
	foo2 = MyStruct{1, "cheese"} // init fields in order
	// all of the must be specified if you do it this way ^
//	foo = MyStruct{foo: 1} // init fields selectively by name
//	foo = 7 // not allowed, becuase 7 doesn't have a Foo function

	fmt.Println("foo", foo.Foo(), "foo2", foo2.Foo())
	// what is an interface?
	// it is a tuple containing an object's type, and instance
	// interfaces correspond to a set of methods(functions) on an object.

}

